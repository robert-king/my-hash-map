package hashmap

import "github.com/draffensperger/golp"

var BitScoreCache [5000][5000]uint16 //constant once written, could be dedicated piece of readonly memory, shared by OS

func diffBit(a, b, i uint64) bool {
	return (a >> i) & 1 != (b >> i) & 1
}

func differentBits(a, b uint64) (ints []uint64) {
	if a < b {
		a,b = b,a
	}
	for i := uint64(0); a >> i > 0; i++ {
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

func bitsToInt(bits []uint16) uint16 {
	bitNum := 0
	for _, bit := range bits {
		bitNum += 1 << bit
	}
	return uint16(bitNum)
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
		[]uint16{0,1},
		[]uint16{0,2},
		[]uint16{5},
		[]uint16{1,2},
		[]uint16{0,3},
		[]uint16{6},
		[]uint16{7},
		[]uint16{8},
		[]uint16{0,4},
		[]uint16{1,3},
		[]uint16{1,4},
		[]uint16{2,3},
		[]uint16{0,5},
		[]uint16{2,4},
		[]uint16{1,5},
		[]uint16{0,1,2},
		[]uint16{0,1,3},
		[]uint16{2,5},
		[]uint16{0,6},
		[]uint16{0,7},
		[]uint16{9},
		[]uint16{0,2,3},
		[]uint16{2,6},
		[]uint16{1,7},
		[]uint16{1,6},
		[]uint16{3,5},
		[]uint16{2,7},
		[]uint16{10},
		[]uint16{1,3,8},
		[]uint16{4,7},
		[]uint16{3,6},
		[]uint16{0,2,11},
		[]uint16{3,8},
		[]uint16{1,2,5},
		[]uint16{0,3,4},
		[]uint16{0,1,12},
		[]uint16{0,1,2,5},
		[]uint16{4,6},
		[]uint16{3,4},
		[]uint16{0,10},
		[]uint16{0,8},
		[]uint16{0,1,4},
		[]uint16{11},
		[]uint16{1,2,3},
		[]uint16{0,2,4},
		[]uint16{1,2,9},
		[]uint16{3,12},
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

	println(bits)

	return bits
}

func addToBitScoreCash(num uint64, bits []uint16, score uint16) {
	bitsNum := bitsToInt(bits)
	matchedBits := bitsNum & uint16(num)
	BitScoreCache[bitsNum][matchedBits] = score
}

func bitScore(num uint64, bits []uint16) (score uint16) {
	for _, bit := range bits {
		if (num >> bit) & 1 == 1 {
			score += 1 << bit
		}
	}
	addToBitScoreCash(num, bits, score)
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

