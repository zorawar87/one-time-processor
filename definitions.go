package main

import (
	"fmt"
	re "regexp"
)

type (
	Text           string
	Raw            []int
	KnownPlaintext struct {
		Cipher Text
		Plain  Text
	}
	KeyPairs map[Text][]KnownPlaintext
)

const (
	CHARSET         = "EHIKLRST"
	DICTIONARY_NAME = "dictionary.txt"
	DICTIONARY_URL  = "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
)

var (
	VALID_WORD = re.MustCompile(fmt.Sprintf("^[%s]+$", CHARSET))
	CIPHERS    = []Text{"KHHLTK", "KTHLLE"}
	// assumes that all ciphertexts are of the same length
	// (which is valid for the given ciphertexts)
	MSG_LENGTH = len(CIPHERS[0])
)
