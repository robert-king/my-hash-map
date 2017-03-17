package algorithms

import (
	"fmt"
	"testing"
	"time"
)

func TestAnagrammars(t *testing.T) {
	words, err := ReadWords()
	if err != nil {
		t.Error(err)
		return
	}
	//words = []string{"abolitionism", "conventionalize"}

	basic := BuildPrimeBasicAnagrammar(words)
	adv := BuildPrimeAdvancedAnagrammar(words)

	anagrammars := []Anagrammar{
		BuildSortAnagrammar(words),
		basic,
		adv,
	}

	for _, ag := range anagrammars {
		if len(ag.GetAnagrams("tar")) != 3 {
			t.Error("invalid anagrams for art")
		}
	}

	fmt.Println(adv.GetAnagrams("datamine"))

	for _, w := range words {
		if len(basic.GetAnagrams(w)) != len(adv.GetAnagrams(w)) {
			fmt.Println(basic.GetAnagrams(w), adv.GetAnagrams(w))
			//t.Error("words not equal")
			//return
		}
	}

	now := time.Now()
	for _, w := range words {
		basic.GetAnagrams(w)
	}
	fmt.Println("bas", time.Since(now))

	now = time.Now()
	for _, w := range words {
		adv.GetAnagrams(w)
	}
	fmt.Println("adv", time.Since(now))

	now = time.Now()
	for _, w := range words {
		basic.GetAnagrams(w)
	}
	fmt.Println("bas", time.Since(now))

	now = time.Now()
	for _, w := range words {
		adv.GetAnagrams(w)
	}
	fmt.Println("adv", time.Since(now))
	now = time.Now()
	for _, w := range words {
		basic.GetAnagrams(w)
	}
	fmt.Println("bas", time.Since(now))

	now = time.Now()
	for _, w := range words {
		adv.GetAnagrams(w)
	}
	fmt.Println("adv", time.Since(now))

}
