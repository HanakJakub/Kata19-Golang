package chainer

import (
	"kata19/dictionary"
	"testing"
)

type TestCase struct {
	startWord string
	endWord   string
	noOfRes   int
	resLength int
}

func TestIt(t *testing.T) {

	testCases := []TestCase{
		TestCase{
			startWord: "ruby",
			endWord:   "code",
			noOfRes:   2,
			resLength: 3,
		},
	}
	var dict dictionary.Dictionary
	var res []Chain

	for _, testCase := range testCases {
		dict = dictionary.LoadDict("../wordlist.txt", len(testCase.startWord))

		res = CreateWordChains(dict, testCase.startWord, testCase.endWord)

		if len(res) != testCase.noOfRes {
			t.Errorf("There should be %d matching chains", testCase.noOfRes)
		}

		for _, r := range res {
			if r.Length() != testCase.resLength {
				t.Errorf("There should be %d items in chain", testCase.resLength)
			}
		}
	}
}
