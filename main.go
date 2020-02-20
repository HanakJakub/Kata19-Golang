package main

import (
	"fmt"
	"kata19/chainer"
	"kata19/dictionary"
	"log"
	"os"
	"time"
)

func main() {
	// start the timer to log execution time
	start := time.Now()

	if len(os.Args) < 3 {
		panic("Start and end word must be set")
	}

	startWord := os.Args[1]
	endWord := os.Args[2]

	// load dictirionary with matching length of startWord
	dict := dictionary.LoadDict("wordlist.txt", len(startWord))

	var res []chainer.Chain

	// check if words are valid
	if isValid(startWord, endWord) {
		res = chainer.CreateWordChains(dict, startWord, endWord)
	} else {
		panic("Start and end words must be equal size")
	}

	// print the results
	fmt.Println("The best results are:")
	for _, c := range res {
		fmt.Println(c.Get())
	}

	elapsed := time.Since(start)
	log.Printf("Executed in %s", elapsed)
}

func isValid(startWord string, endWord string) bool {
	return len(startWord) == len(endWord)
}
