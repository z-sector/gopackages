package my_random

import (
	"github.com/z-sector/gopackages/wordz"
	"strings"
)

var digits []string

func init() {
	for _, value := range wordz.Words {
		digits = append(digits, strings.ToLower(value))
	}
}

func Digit() string {
	old := wordz.Words
	defer func() {
		wordz.Words = old
	}()
	wordz.Words = digits
	return wordz.Random()
}
