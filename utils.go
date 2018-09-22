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
	for _, char := range strings.Split(ToUpper(string(t)), "") {
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

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

/* Printing Utilties */
func displayHeader() {
	fmt.Print("+=====+========+========+========+\n")
	fmt.Print("|  #  | Cipher | Plain  |  Key   |")
	fmt.Print(" Proof: Cipher \u2295 Plain = Key\n")
	fmt.Print("+=====+========+========+========+\n")
}

func display(count int, cipher Text, plain Text, key Raw) {
	fmt.Printf(
		"| %3d | %-6s | %-6s | %-6s | %s \u2295 %s = %s\n",
		count,
		cipher, plain, key.Decode(),
		fmt.Sprintf("%03b", cipher.Encode()),
		fmt.Sprintf("%03b", plain.Encode()),
		fmt.Sprintf("%03b", key),
	)
}

func displayFooter() {
	fmt.Print("+=====+========+========+========+\n")
}
