package main

import (
	"fmt"
	re "regexp"
)

type Text string
type Raw []int

type Observation struct {
	CipherText Text
	PlainText  Text
	Key        Text
}

const (
	FILENAME = "words.txt"
	CHARSET  = "EHIKLRST"
)

var VALID_WORD = re.MustCompile(fmt.Sprintf("^[%s]+$", CHARSET))
var CIPHERS = []Text{"KHHLTK", "KTHLLE"}

//KHHLTK
//KTHLLE
//|x||xx
// Find words in which
//0,2,3 letters are the same
// for some given key
