package main

import (
	"fmt"
	"github.com/robert-king/fast-anagram/algorithms"
	"time"
	"github.com/pkg/profile"
)

var now time.Time

func init() {
	now = time.Now()
}

func p(s string) {
	fmt.Println(s, time.Since(now))
	now = time.Now()
}

func main() {
	defer profile.Start().Stop()
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

	for i := 0; i < 50; i++ {
		for _, w := range words {
			adv.GetAnagrams(w)
		}
	}

	p("adv")

	for i := 0; i < 50; i++ {
		for _, w := range words {
			basic.GetAnagrams(w)
		}
	}

	p("basic")
}
