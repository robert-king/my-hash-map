package hashmap

const numOnes = uint8(20)
const ones = uint32(1<<numOnes - 1)

type Info struct {
	bitNum uint16
	start  uint32
}

type Map struct {
	info [ones + 1]Info
}

func NewMap(hashes []uint64) (*Map, uint32) {
	var collisions [ones + 1][]uint64
	var info [ones + 1]Info

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
		info[part].bitNum = bitsToInt(bits)
		info[part].start = start
		start += uint32(maxBitScore(nums, bits)) + 1
	}
	locator := &Map{
		info: info,
	}
	return locator, start
}

func (m *Map) Index(num uint64) uint32 {
	part := uint32(num) & ones
	info := m.info[part]
	if info.bitNum == 0 {
		return info.start
	}

	remaining := num >> numOnes
	matchedBits := info.bitNum & uint16(remaining)

	if matchedBits == 0 {
		return info.start
	}

	//if info.bitNum == 1 || info.bitNum == 2 || info.bitNum == 4 || info.bitNum == 8 { //info.bitNum & (info.bitNum - 1)) == 0
	//	return info.start + uint32(1)
	//}

	offset := BitScoreCache[matchedBits*4100+info.bitNum]
	return info.start + uint32(offset)
}
