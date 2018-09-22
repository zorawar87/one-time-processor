package main

import (
	"bufio"
	"os"
)

// For each ciphertext
// candidate plaintexts are obtained
// for which an associated OTP key is deduced and displayed
func main() {
	for _, cipher := range CIPHERS {
		displayHeader()
		words := make(chan Text)

		go getPlaintext(len(cipher), words)
		getKeys(cipher, words)

		displayFooter()
	}
}

// From a local word list,
// candidate words are retrieved based on
// its length and its constiuting characters
func getPlaintext(lenCipher int, words chan Text) {
	var err error
	file, err := os.Open(FILENAME)
	panicOnErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := ToUpper(scanner.Text())
		if len(word) == lenCipher && VALID_WORD.MatchString(word) {
			words <- Text(word)
		}
	}
	close(words)
}

// For each candidate plaintext word,
// an OTP key is deduced by XOR-ing the plaintext with the cipher
func getKeys(cipher Text, words chan Text) {
	lenCipher := len(cipher)
	wordCount := 0

	var key Raw
	for word := range words {
		key = make([]int, 0, lenCipher)
		wordCount++

		for j := 0; j < lenCipher; j++ {
			key = append(key, cipher.Encode()[j]|word.Encode()[j])
		}

		display(wordCount, cipher, word, key)
	}
}
