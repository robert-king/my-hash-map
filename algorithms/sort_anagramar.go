package algorithms

type SortAnagrammar struct {
	m map[string][]string
}

func BuildSortAnagrammar(words []string) SortAnagrammar {
	sa := SortAnagrammar{
		make(map[string][]string),
	}
	for _, word := range words {
		h := SortString(word)
		sa.m[h] = append(sa.m[h], word)
	}
	return sa
}

func (sa SortAnagrammar) GetAnagrams(word string) []string {
	return sa.m[SortString(word)]
}
