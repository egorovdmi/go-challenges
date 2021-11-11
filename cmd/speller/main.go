package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/egorovdmi/go-challenges/internals/platform/dictionaries"
)

const DEFAULT_DICTIONARY_FILE = "./assets/speller-hash/dictionaries/large"

type WordsHandler func(w string)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: speller [dictionary] text")
		return
	}

	// Load command line agruments =========================================
	var cfg struct {
		dictionaryFile string
		textFile       string
	}

	if len(args) > 1 {
		cfg.dictionaryFile = args[0]
		cfg.textFile = args[1]
	} else {
		cfg.dictionaryFile = DEFAULT_DICTIONARY_FILE
		cfg.textFile = args[0]
	}

	d := dictionaries.NewDictionary()

	// Load words into dictionary
	readWordsFromFile(cfg.dictionaryFile, func(w string) {
		d.Add(w)
	})

	// Validate a text
	fmt.Println("MISSPELLED WORDS:")
	missspelled := 0
	totalWords := 0

	readWordsFromFile(cfg.textFile, func(w string) {
		totalWords++
		if !d.Contains(w) {
			missspelled++
			fmt.Printf("'%v'\n", w)
		}
	})

	// Print statistics
	fmt.Printf("WORDS MISSPELLED: %v\n", missspelled)
	fmt.Printf("WORDS IN DICTIONARY: %v\n", d.Count)
	fmt.Printf("LISTS IN DICTIONARY: %v\n", d.ListsCount)
	fmt.Printf("WORDS IN TEXT: %v\n", totalWords)
}

func readWordsFromFile(filePath string, wordsHandler WordsHandler) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		s := scanner.Text()
		wordsHandler(s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
