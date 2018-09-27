package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Encodes text to its binary equivalent,
// based on the provided (character<=>binary) spec.
// Undoes (r Raw)Decode() Text
func (t Text) Encode() (binary Raw) {
	for _, char := range strings.Split(strings.ToUpper(string(t)), "") {
		binary = append(binary, strings.Index(CHARSET, char))
	}
	return
}

// Decodes binary to its text equivalent,
// based on the provided (character<=>binary) spec.
// Undoes (t Text)Encode() Raw
func (r Raw) Decode() Text {
	var b bytes.Buffer
	for _, v := range r {
		b.WriteString(fmt.Sprintf("%c", CHARSET[v]))
	}
	return Text(b.String())
}

// XORs the receiver with the parameter
// Note:
//     t      |  u          |   result
// Plaintext  | Key        ==> Ciphertext
// Ciphertext | Key        ==> Plaintext
// Plaintext  | Ciphertext ==> Key
func (t Text) XorEachChar(u Text) Text {
	result := make(Raw, 0, len(t))
	for i := 0; i < len(t); i++ {
		result = append(result, t.Encode()[i]^u.Encode()[i])
	}
	return result.Decode()
}
