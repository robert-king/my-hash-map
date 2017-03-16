package algorithms

import (
	"github.com/robert-king/fast-anagram/hashmap"
)

type PrimeAdvancedAnagrammar struct {
	m [][]string
	locator *hashmap.Map
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

	locator, maxIdx := hashmap.NewMap(hashes)
	grams := make([][]string, maxIdx)
	for _, word := range words {
		idx := locator.Index(PrimeProduct(word))
		grams[idx] = append(grams[idx], word)
	}
	println(len(words))
	println(len(grams))
	return PrimeAdvancedAnagrammar{
		m: grams,
		locator: locator,
	}
}

func (pa PrimeAdvancedAnagrammar) GetAnagrams(word string) []string {
	return pa.m[pa.locator.Index(PrimeProduct(word))]
}