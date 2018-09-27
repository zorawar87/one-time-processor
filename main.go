package main

//
// 1. Retrieves valid plaintext from a local dictionary
//		"valid" is determined by the word's
// 			- length (equal to length of cipher)
// 			- charset (as specified)
// 2. Populates possible key-(plaintext-ciphertext) pairs
// 		by computing "ciphertext XOR plaintext" to produce key
// 3. Displays all keys and ciphertexts which produce
//    a valid decryption, i.e., both ciphertexts yield
//    valid plaintext with the same key
func main() {
	words := make(chan text)
	pairs := make(keyPairs)

	go getPlaintext(words)
	pairs.populate(words)
	pairs.displayValidDecryptions()
}
