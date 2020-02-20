package dictionary

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Dictionary contains all words that will be used for searching chain
type Dictionary struct {
	// Words is a list of words that have exact same length as start word of chain
	Words []string
}

// LoadDict will load dictionary from a file and return a dictionary struct
func LoadDict(path string, length int) (d Dictionary) {
	words, err := readLines(path, length)
	if err != nil {
		panic(err)
	}

	d.Words = words

	return
}

// filterWords will filter words based on compareWord
func (d Dictionary) filterWords(startWord string, endWord string, pos int) map[string]string {
	// Clear the filter list
	filtered := make(map[string]string)

	for i := 0; i < len(d.Words); i++ {
		if len(d.Words[i]) == len(startWord) {
			if shouldAppendToList(startWord, endWord, d.Words[i], pos) {
				filtered[d.Words[i]] = d.Words[i]
			}
		}
	}

	return filtered
}

// GetFilteredWords will return filtered words
// if there are no words matching on position for startword and endword
// it will try to filter words only for startword
func (d Dictionary) GetFilteredWords(startWord string, endWord string, pos int) map[string]string {
	// Filter out with start and end words
	filtered := d.filterWords(startWord, endWord, pos)

	// If empty list try it without matching endword
	if len(filtered) == 0 {
		filtered = d.filterWords(startWord, "", pos)
		pos--
	}

	return filtered
}

// WordScore will calculate score of similarity between word and toCompare
func WordScore(word string, toCompare string) int {
	score := 0

	for i := 0; i < len(word); i++ {
		if word[i] == toCompare[i] {
			score++
		}
	}

	return score
}

func readLines(path string, length int) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			line := buffer.String()

			// filter out lines that are longer that length
			if len(line) == length {
				lines = append(lines, buffer.String())
			}
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func shouldAppendToList(startWord string, endWord string, word string, pos int) bool {
	startWordScoreHasOneCharDiff := WordScore(startWord, word) == len(startWord)-1
	endWordScoreHasAtLeastPosCharSame := WordScore(endWord, word) >= pos

	return startWordScoreHasOneCharDiff && (endWord == "" || endWordScoreHasAtLeastPosCharSame)
}
