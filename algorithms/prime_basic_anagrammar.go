package algorithms


type PrimeBasicAnagrammar struct {
	m map[uint64][]string
}

func BuildPrimeBasicAnagrammar(words []string) PrimeBasicAnagrammar {
	pba := PrimeBasicAnagrammar{
		make(map[uint64][]string),
	}
	for _, word := range words {
		h := PrimeProduct(word)
		pba.m[h] = append(pba.m[h], word)
	}
	println(len(pba.m), "!!")
	return pba
}

func (pba PrimeBasicAnagrammar) GetAnagrams(word string) []string {
	return pba.m[PrimeProduct(word)]
}