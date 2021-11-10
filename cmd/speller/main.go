package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/egorovdmi/go-challenges/internals/platform/dictionaries"
)

const DEFAULT_DICTIONARY_FILE = "./assets/speller-hash/dictionaries/large"

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

	stringsForDict := readStringsFromFile(cfg.dictionaryFile)
	textStrings := readStringsFromFile(cfg.textFile)

	d := dictionaries.NewDictionary()
	d.Load(stringsForDict)

	fmt.Println("MISSPELLED WORDS:")
	missspelled := 0
	for _, w := range textStrings {
		if !d.Contains(w) {
			missspelled++
			fmt.Printf("'%v'\n", w)
		}
	}

	fmt.Printf("WORDS MISSPELLED: %v\n", missspelled)
	fmt.Printf("WORDS IN DICTIONARY: %v\n", d.Count)
	fmt.Printf("LISTS IN DICTIONARY: %v\n", d.ListsCount)
	fmt.Printf("WORDS IN TEXT: %v\n", len(textStrings))
}

func readStringsFromFile(filePath string) []string {
	var result []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		s := scanner.Text()
		result = append(result, s)
		// fmt.Printf("%v has scanned.\n", s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
