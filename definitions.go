package main

import (
	"fmt"
	re "regexp"
)

type (
	// text can represent CipherText, PlainText or the Key (as text)
	text string
	// raw can represent any Text variable as raw bytes, decoded per the charset definition
	raw            []int
	knownPlaintext struct {
		Cipher text
		Plain  text
	}
	keyPairs map[text][]knownPlaintext
)

const (
	charset        = "EHIKLRST"
	dictionaryName = "dictionary.txt"
	dictionaryURL  = "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
)

var (
	validWord = re.MustCompile(fmt.Sprintf("^[%s]+$", charset))
	ciphers   = []text{"KHHLTK", "KTHLLE"}
	// assumes that all ciphertexts are of the same length
	// (which is valid for the given ciphertexts)
	messageLength = len(ciphers[0])
)
