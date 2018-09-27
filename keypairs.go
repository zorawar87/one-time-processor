package main

import "fmt"

// Populate the receiver with all possible key-(ciphertext,plaintext) pairs
func (pairs keyPairs) populate(words chan text) {
	for plaintext := range words {
		for _, cipher := range ciphers {
			// ciphertext XOR plaintext produces the key
			key := cipher.XorEachChar(plaintext)
			pairs[key] = append(pairs[key], knownPlaintext{cipher, plaintext})
		}
	}
}

// displays all valid decryptions
func (pairs keyPairs) displayValidDecryptions() {
	fmt.Print("+======+========+========+========+\n")
	fmt.Print("| Msg# |  Key   | Cipher | Plain  |\n")
	fmt.Print("+======+========+========+========+\n")
	messageCount := 0
	for key, knownPlaintexts := range pairs {
		if len(knownPlaintexts) == len(ciphers) {
			messageCount++
			for _, v := range knownPlaintexts {
				fmt.Printf(
					"| %3d  | %-6s | %-6s | %-6s |\n",
					messageCount, key, v.Cipher, v.Plain,
				)
			}
		}
	}
	fmt.Print("+======+========+========+========+\n")
}
