package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Candidate words are chosen from a local dictionary based on:
//  - its length
//	- its charset (must match)
// Accepted candidates are piped through the `words` channel
// for further processing
func getPlaintext(words chan Text) {
	var err error

	DownloadDictionaryIfNotExists()

	file, err := os.Open(DICTIONARY_NAME)
	panicOnErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToUpper(scanner.Text())
		if len(word) == MSG_LENGTH && VALID_WORD.MatchString(word) {
			words <- Text(word)
		}
	}
	close(words)
}

// Downloads a remote dictionary word list (<5MB)
// if a dictionary does not exist locally
func DownloadDictionaryIfNotExists() {
	if _, err := os.Stat(DICTIONARY_NAME); !os.IsNotExist(err) {
		fmt.Println("Shaved 5MB")
		return
	}

	// create a local storage file
	out, err := os.Create(DICTIONARY_NAME)
	panicOnErr(err)
	defer out.Close()

	// retrieve the dictionary
	resp, err := http.Get(DICTIONARY_URL)
	panicOnErr(err)
	defer resp.Body.Close()

	// Write the dictionary to file
	_, err = io.Copy(out, resp.Body)
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
