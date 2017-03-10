package algorithms

import (
	"unsafe"
)

type PrimeAdvancedAnagrammar struct {
	m [][]string
	locator *Locator
}

func BuildPrimeAdvancedAnagrammar(words []string) PrimeAdvancedAnagrammar {
	hashesMap := make(map[uint64]bool,0)
	for _, word := range words {
		hashesMap[PrimeProduct(word)] = true
	}
	hashes := make([]uint64, 0)
	for hash, _ := range hashesMap {
		hashes = append(hashes, hash)
	}

	locator := NewLocator(hashes)
	grams := make([][]string, locator.maxIndex)
	for _, word := range words {
		idx := locator.Index1(PrimeProduct(word))
		grams[idx] = append(grams[idx], word)
	}
	start := uintptr(unsafe.Pointer(&grams[0]))
	size := unsafe.Sizeof(grams[0])
	locator.UpdateStarts(start, size)
	println(len(words))
	println(len(grams))
	println(len(locator.starts))
	return PrimeAdvancedAnagrammar{
		m: grams,
		locator: locator,
	}
}

func (pa PrimeAdvancedAnagrammar) GetAnagrams(word string) []string {
	return *(*[]string)(unsafe.Pointer(pa.locator.Index(PrimeProduct(word))))
}