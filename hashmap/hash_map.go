package hashmap

const numOnes = uint8(20)
const ones = uint32(1 << numOnes - 1)



type Map struct {
	bitNums [ones+1]uint16
	starts [ones+1]uint32
}

func NewMap(hashes []uint64) (*Map, uint32) {
	var collisions [ones+1][]uint64
	var bitNums [ones+1]uint16
	var starts [ones+1]uint32

	for _, hash := range hashes {
		part := uint32(hash) & ones
		remaining := hash >> numOnes
		collisions[part] = append(collisions[part], remaining)
	}
	start := uint32(1)
	for part, nums := range collisions {
		if len(nums) == 0 {
			continue
		}
		bits := minimumDistinguishingBits(nums)
		bitNums[part] = bitsToInt(bits)
		starts[part] = start
		start += uint32(maxBitScore(nums, bits)) + 1
	}
	locator := &Map{
		bitNums: bitNums,
		starts: starts,
	}
	return locator, start
}

func (m *Map) Index(num uint64) uint32 {
	part := uint32(num) & ones
	remaining := num >> numOnes
	start := m.starts[part]
	bitsNum := m.bitNums[part]
	matchedBits := bitsNum & uint16(remaining)
	offset := BitScoreCache[bitsNum][matchedBits]
	return start + uint32(offset)
}