package main

import "fmt"

//Trie Node

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

//Trie
type Trie struct {
	root *TrieNode
}

// init function

// func initTrie() *Trie {
// 	myTrie := &Trie{}
// 	myTrie.root = &TrieNode{}
// 	return myTrie
// }

func (t *Trie) insert(w string) {
	wordLength := len(w)
	newNode := new(TrieNode)
	temp := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
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
		fmt.Println(w[i])
		if temp.children[charIndex] == nil {
			return false
		}
		temp = temp.children[charIndex]
		//fmt.Println(temp.children[i])
	}
	if temp.isEnd == true {
		return true
	}
	return false

}

func main() {
	SampleTrie := Trie{}
	SampleTrie.root = &TrieNode{}
	SampleTrie.insert("sanika")
	SampleTrie.insert("jerin")
	SampleTrie.insert("saniha")
	fmt.Println(SampleTrie.Search("sanika"))
	fmt.Println(SampleTrie.Search("sanniha"))

}
