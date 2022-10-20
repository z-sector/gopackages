package my_random

import "github.com/z-sector/gopackages/wordz"

var cities = []string{"City1", "City2", "City3"}

func City() string {
	old := wordz.Words
	defer func() {
		wordz.Words = old
	}()
	wordz.Words = cities
	return wordz.Random()
}
