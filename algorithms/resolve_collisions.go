package algorithms

import (
	"fmt"
	"github.com/draffensperger/golp"
)

const numOnes = uint64(20)
const ones = uint64(1<<numOnes - 1)

func diffBit(a, b, i uint64) bool {
	return (a>>i)&1 != (b>>i)&1
}

func differentBits(a, b uint64) (ints []uint64) {
	if a < b {
		a, b = b, a
	}
	for i := uint64(0); a>>i > 0; i++ {
		if diffBit(a, b, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

func checkBitsDistinguishNums(nums []uint64, bits []uint16) bool {
	offsets := make(map[uint16]bool)
	for _, num := range nums {
		offset := bitScore(num, bits)
		_, ok := offsets[offset]
		if ok {
			return false
		}
		offsets[offset] = true
	}
	return true
}

func minimumDistinguishingBits(nums []uint64) (bits []uint16) {
	if len(nums) == 1 {
		return bits
	}

	var likelySolutions = [][]uint16{
		[]uint16{0},
		[]uint16{1},
		[]uint16{2},
		[]uint16{3},
		[]uint16{4},
		[]uint16{0, 1},
		[]uint16{0, 2},
		[]uint16{5},
		[]uint16{1, 2},
		[]uint16{0, 3},
		[]uint16{6},
		[]uint16{7},
		[]uint16{8},
		[]uint16{0, 4},
		[]uint16{1, 3},
		[]uint16{1, 4},
		[]uint16{2, 3},
		[]uint16{0, 5},
		[]uint16{2, 4},
		[]uint16{1, 5},
		[]uint16{0, 1, 2},
		[]uint16{0, 1, 3},
		[]uint16{2, 5},
		[]uint16{0, 6},
		[]uint16{0, 7},
		[]uint16{9},
		[]uint16{0, 2, 3},
		[]uint16{2, 6},
		[]uint16{1, 7},
		[]uint16{1, 6},
		[]uint16{3, 5},
		[]uint16{2, 7},
		[]uint16{10},
		[]uint16{1, 3, 8},
		[]uint16{4, 7},
		[]uint16{3, 6},
		[]uint16{0, 2, 11},
		[]uint16{3, 8},
		[]uint16{1, 2, 5},
		[]uint16{0, 3, 4},
		[]uint16{0, 1, 12},
		[]uint16{0, 1, 2, 5},
		[]uint16{4, 6},
		[]uint16{3, 4},
		[]uint16{0, 10},
		[]uint16{0, 8},
		[]uint16{0, 1, 4},
		[]uint16{11},
		[]uint16{1, 2, 3},
		[]uint16{0, 2, 4},
		[]uint16{1, 2, 9},
		[]uint16{3, 12},
	}

	for _, likelySolution := range likelySolutions {
		if checkBitsDistinguishNums(nums, likelySolution) {
			return likelySolution
		}
	}

	var equations [][]golp.Entry
	maxDiffBit := uint64(0)
	for i := range nums {
		for j := 0; j < i; j++ {
			var equation []golp.Entry
			for _, bit := range differentBits(nums[i], nums[j]) {
				if bit > maxDiffBit {
					maxDiffBit = bit
				}
				equation = append(equation, golp.Entry{int(bit), 1})
			}
			equations = append(equations, equation)
		}
	}
	lp := golp.NewLP(0, int(maxDiffBit)+1)
	for _, equation := range equations {
		lp.AddConstraintSparse(equation, golp.GE, 1)
	}
	var objFn []float64
	for i := 0; i <= int(maxDiffBit); i++ {
		lp.SetInt(i, true)
		objFn = append(objFn, float64(i))
	}
	lp.SetObjFn(objFn)
	lp.Solve()
	for i, bit := range lp.Variables() {
		if bit == 1 {
			bits = append(bits, uint16(i))
		}
	}

	fmt.Println("Likely solution not found but found ", bits)

	return bits
}

func bitScore(num uint64, bits []uint16) (score uint16) {
	for _, bit := range bits {
		if (num>>bit)&1 == 1 {
			score += 1 << bit
		}
	}
	return score
}

func maxBitScore(nums []uint64, bits []uint16) uint16 {
	mx := uint16(0)
	for _, num := range nums {
		bs := bitScore(num, bits)
		if bs > mx {
			mx = bs
		}
	}
	return mx
}

type Locator struct {
	collisions [ones + 1][]uint64
	bits       [ones + 1][]uint16
	starts     [ones + 1]uintptr
	maxIndex   uintptr
}

func NewLocator(hashes []uint64) *Locator {
	var collisions [ones + 1][]uint64
	var bits [ones + 1][]uint16
	var starts [ones + 1]uintptr

	for _, hash := range hashes {
		part := hash & ones
		remaining := hash >> numOnes
		collisions[part] = append(collisions[part], remaining)
	}
	start := uintptr(1)
	for part, nums := range collisions {
		if len(nums) == 0 {
			continue
		}
		bits[part] = minimumDistinguishingBits(nums)
		if !checkBitsDistinguishNums(nums, bits[part]) {
			fmt.Println("ERROR!")
		}
		//if len(nums) > 1 {
		//	fmt.Println("nums", nums)
		//	fmt.Println("bits", bits[part])
		//}
		starts[part] = start
		start += uintptr(maxBitScore(nums, bits[part])) + 1
	}
	locator := &Locator{
		collisions: collisions,
		bits:       bits,
		starts:     starts,
		maxIndex:   start,
	}
	return locator
}

func (locator *Locator) UpdateStarts(start, size uintptr) {
	for i := range locator.starts {
		if locator.starts[i] > 0 {
			locator.starts[i] = start + locator.starts[i]*size
		}
	}
}

func (locator *Locator) Index(num uint64) uintptr {
	part := num & ones
	remaining := num >> numOnes
	start := locator.starts[part]
	bits := locator.bits[part]
	offset := 24 * uintptr(bitScore(remaining, bits))
	return start + offset
}

func (locator *Locator) Index1(num uint64) uintptr {
	part := num & ones
	//remaining := num >> numOnes
	//start := locator.starts[part]
	//bits := locator.bits[part]
	//offset := bitScore(num >> numOnes, locator.bits[part])
	if len(locator.collisions[part]) == 1 {
		return locator.starts[part]
	}
	return locator.starts[part] + uintptr(bitScore(num>>numOnes, locator.bits[part]))
}
