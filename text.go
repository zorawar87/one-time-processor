package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Encodes text to its binary equivalent,
// based on the provided (character<=>binary) spec.
// Undoes (r raw)Decode() text
func (t text) Encode() (binary raw) {
	for _, char := range strings.Split(strings.ToUpper(string(t)), "") {
		binary = append(binary, strings.Index(charset, char))
	}
	return
}

// Decodes binary to its text equivalent,
// based on the provided (character<=>binary) spec.
// Undoes (t text)Encode() raw
func (r raw) Decode() text {
	var b bytes.Buffer
	for _, v := range r {
		b.WriteString(fmt.Sprintf("%c", charset[v]))
	}
	return text(b.String())
}

// XORs the receiver with the parameter
// Note:
//     t      |  u          |   result
// Plaintext  | Key        ==> Ciphertext
// Ciphertext | Key        ==> Plaintext
// Plaintext  | Ciphertext ==> Key
func (t text) XorEachChar(u text) text {
	result := make(raw, 0, len(t))
	for i := 0; i < len(t); i++ {
		result = append(result, t.Encode()[i]^u.Encode()[i])
	}
	return result.Decode()
}
