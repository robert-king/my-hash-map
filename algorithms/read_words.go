package algorithms

import (
	"io/ioutil"
	"strings"
)

func ReadWords() ([]string, error) {
	data, err := ioutil.ReadFile("dict.txt")
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(data),"\n")
	return words, nil
}