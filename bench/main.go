package main

import (
	"fmt"
	"github.com/pkg/profile"
	"github.com/robert-king/fast-anagram/algorithms"
	"time"
)

var now time.Time

func init() {
	now = time.Now()
}

func p(s string) {
	fmt.Println(s, time.Since(now))
	now = time.Now()
}

func runWithAnagrams() {
	words, err := algorithms.ReadWords()
	if err != nil {
		fmt.Println("error reading words")
		return
	}
	//words = []string{"abolitionism", "conventionalize"}
	p("read words")

	basic := algorithms.BuildPrimeBasicAnagrammar(words)
	p("basic built")
	adv := algorithms.BuildPrimeAdvancedAnagrammar(words)
	p("advanced built")

	for _, w := range words {
		if len(basic.GetAnagrams(w)) != len(adv.GetAnagrams(w)) {
			fmt.Println(basic.GetAnagrams(w), adv.GetAnagrams(w))
			//t.Error("words not equal")
			//return
		}
	}
	p("checked valid")

	tests := 500

	for _, w := range words {
		for i := 0; i < tests; i++ {
			adv.GetAnagrams(w)
		}
	}

	p("adv")

	for _, w := range words {
		for i := 0; i < tests; i++ {
			basic.GetAnagrams(w)
		}
	}

	p("basic")

	for i := 0; i < tests; i++ {
		for _, w := range words {
			adv.GetAnagrams(w)
		}
	}

	p("adv")

	for i := 0; i < tests; i++ {
		for _, w := range words {
			basic.GetAnagrams(w)
		}
	}

	p("basic")
}

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	runWithAnagrams()
	//adv 5.280358926s
	//basic 6.118879895s
	//adv 19.165513092s
	//basic 16.734885256s

}
