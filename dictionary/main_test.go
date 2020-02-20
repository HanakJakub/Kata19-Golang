package dictionary

import (
	"testing"
)

func TestLoadDictionary(t *testing.T) {
	dict := LoadDict("../wordlist.txt", 3)

	for _, word := range dict.Words {
		if len(word) != 3 {
			t.Error("In dictionary words should be only words of length 3")
		}
	}
}

func TestWordScoreShouldReturnNumberOfMatchedChars(t *testing.T) {
	a := "testing"
	b := "testing"

	if WordScore(a, b) != 7 {
		t.Error("It should return that all chars are the same")
	}

	a = "loading"
	b = "testing"

	if WordScore(a, b) != 3 {
		t.Error("It should matched ing in the end")
	}
}

func TestGetFilteredWordsShouldReturnWordsWithMatchingScores(t *testing.T) {
	dict := LoadDict("../wordlist.txt", 3)

	filtered := dict.GetFilteredWords("cat", "dog", 1)

	if len(filtered) != 2 {
		t.Error("It should return only 2 matched words with one change in word cat and 1 match in word dog")
	}

	dict = LoadDict("../wordlist.txt", 4)

	filtered = dict.GetFilteredWords("lead", "code", 1)

	if len(filtered) != 1 {
		t.Error("It should return only 1 matched word with one chage in word lead and one match in code")
	}
}
