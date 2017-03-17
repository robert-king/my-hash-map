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

	for _, w := range words {
		for i := 0; i < 50; i++ {
			adv.GetAnagrams(w)
		}
	}

	p("adv")

	for _, w := range words {
		for i := 0; i < 50; i++ {
			basic.GetAnagrams(w)
		}
	}

	p("basic")
}

func main() {
	defer profile.Start().Stop()
	runWithAnagrams()
	//adv 2.210262363s
	//basic 1.681843672s

}
