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
func getPlaintext(words chan text) {
	var err error

	downloadDictionaryIfNotExists()

	file, err := os.Open(dictionaryName)
	panicOnErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToUpper(scanner.Text())
		if len(word) == messageLength && validWord.MatchString(word) {
			words <- text(word)
		}
	}
	close(words)
}

// Downloads a remote dictionary word list (<5MB)
// if a dictionary does not exist locally
func downloadDictionaryIfNotExists() {
	if _, err := os.Stat(dictionaryName); !os.IsNotExist(err) {
		fmt.Println("Shaved 5MB")
		return
	}

	// create a local storage file
	out, err := os.Create(dictionaryName)
	panicOnErr(err)
	defer out.Close()

	// retrieve the dictionary
	resp, err := http.Get(dictionaryURL)
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
