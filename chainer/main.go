package chainer

import (
	"kata19/dictionary"
	"log"
	"sync"
)

var (
	min    = 9999
	chains = make(chan Chain)
	wg     = new(sync.WaitGroup)
)

// Chain struct that contains list
type Chain struct {
	list      []string
	StartWord string
	EndWord   string
}

// Append to the list of chain
func (c *Chain) Append(s ...string) {
	c.list = append(c.list, s...)
}

// Length will return length of list
func (c *Chain) Length() int {
	return len(c.list)
}

// Get will return the prepared full chain
func (c *Chain) Get() []string {
	prepared := []string{c.StartWord}

	prepared = append(prepared, c.list...)
	prepared = append(prepared, c.EndWord)

	return prepared
}

// GetList will return the full list
func (c *Chain) GetList() []string {
	return c.list
}

// Contains will search if chain list contains given needle
func (c *Chain) Contains(needle string) bool {
	if needle == c.StartWord || needle == c.EndWord {
		return true
	}

	for i := 0; i < len(c.list); i++ {
		if c.list[i] == needle {
			return true
		}
	}

	return false
}

// CreateWordChains will create all possible word chains
func CreateWordChains(dict dictionary.Dictionary, startWord string, endWord string) []Chain {
	var ch Chain

	for _, f := range dict.GetFilteredWords(startWord, endWord, 1) {
		wg.Add(1)
		ch = Chain{StartWord: startWord, EndWord: endWord}

		// start chaining
		go chaining(ch, f, dict, 1)
	}

	// wait untill the chaining process is finished
	go func(wg *sync.WaitGroup) {
		log.Println("waiting")
		wg.Wait()
		close(chains)
	}(wg)

	// process chaining results
	var res []Chain
	for i := range chains {
		res = append(res, i)
	}

	return getOnlyBestResults(res)
}

func chaining(chain Chain, word string, dict dictionary.Dictionary, pos int) {
	defer wg.Done()

	if word != chain.EndWord && pos < len(chain.StartWord) {
		if !chain.Contains(word) {
			chain.Append(word)
			newCh := newChain(chain)
			newCh.Append(chain.GetList()...)

			for _, f := range dict.GetFilteredWords(word, chain.EndWord, pos) {
				wg.Add(1)
				go chaining(newCh, f, dict, pos+1)
			}
		}
	}

	if dictionary.WordScore(word, chain.EndWord) == len(chain.EndWord)-1 {
		if !chain.Contains(word) {
			chain.Append(word)
		}

		if min >= chain.Length() {
			min = chain.Length()

			chains <- chain
		}
	}
}

func newChain(chain Chain) Chain {
	ch := Chain{}

	ch.StartWord = chain.StartWord
	ch.EndWord = chain.EndWord

	return ch
}

func getOnlyBestResults(chains []Chain) (onlyMinChains []Chain) {
	for _, c := range chains {
		if c.Length() == min {
			onlyMinChains = append(onlyMinChains, c)
		}
	}

	return
}
