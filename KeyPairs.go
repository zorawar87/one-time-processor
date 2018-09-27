package main

import "fmt"

// Populate the receiver with all possible key-(ciphertext,plaintext) pairs
func (m KeyPairs) populate(words chan Text) {
	for plaintext := range words {
		for _, cipher := range CIPHERS {
			// ciphertext XOR plaintext produces the key
			key := cipher.XorEachChar(plaintext)
			m[key] = append(m[key], KnownPlaintext{cipher, plaintext})
		}
	}
}

// displays all valid decryptions
func (keyPairs KeyPairs) displayValidDecryptions() {
	fmt.Print("+======+========+========+========+\n")
	fmt.Print("| Msg# |  Key   | Cipher | Plain  |\n")
	fmt.Print("+======+========+========+========+\n")
	messageCount := 0
	for key, knownPlaintexts := range keyPairs {
		if len(knownPlaintexts) == len(CIPHERS) {
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
