package main

import (
	"fmt"
	re "regexp"
)

type Text string
type Raw []int

const (
	FILENAME = "words.txt"
	CHARSET  = "EHIKLRST"
)

var VALID_WORD = re.MustCompile(fmt.Sprintf("^[%s]+$", CHARSET))
var CIPHERS = []Text{"KHHLTK", "KTHLLE"}
