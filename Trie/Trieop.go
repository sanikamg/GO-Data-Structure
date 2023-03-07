package main

import "fmt"

const size = 26

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(w string) {
	wordLength := len(w)
	temp := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			newNode := new(TrieNode)
			temp.children[charIndex] = newNode
		}
		temp = temp.children[charIndex]
	}
	temp.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	temp := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			return false
		}
		temp = temp.children[charIndex]
	}
	if temp.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := Trie{}
	myTrie.root = &TrieNode{}
	myTrie.insert("sanika")
	myTrie.insert("sandhya")
	myTrie.insert("athira")
	fmt.Println(myTrie.Search("sannika"))

}
