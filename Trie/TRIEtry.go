package main

import "fmt"

const size = 26

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

type Trie struct {
	root *trieNode
}

func (t *Trie) buildTrie(w string) {
	wordLength := len(w)
	temp := t.root
	for i := 0; i < wordLength; i++ {
		index := w[i] - 'a'
		if temp.children[index] == nil {
			newNode := new(trieNode)
			temp.children[index] = newNode
		}
		temp = temp.children[index]
	}
	temp.isEnd = true
}

func (t *trieNode) PrintTrie(prefix string) {
	if t.isEnd {
		fmt.Println(prefix)
	}
	for i, child := range t.children {
		if child != nil {
			child.PrintTrie(prefix + string('a'+i))
		}
	}
}

// func (t *Trie) search_by_Prefix(node *trieNode, w string) {
// 	if node.isEnd {
// 		fmt.Println(w)
// 	}
// 	index := w[len(w)-1] - 'a'
// 	if node.children[index] != nil {
// 		new_letter := w + string(index+'a')
// 		t.search_by_Prefix(node.children[index], new_letter)
// 	}

// }

func (t *Trie) searchByPrefix(node *trieNode, word string) {
	if node.isEnd {
		fmt.Println(word)
	}
	index := word[len(word)-1] - 'a'
	if child := node.children[index]; child != nil {
		newWord := word + string(index+'a')
		t.searchByPrefix(child, newWord)
	}
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	temp := t.root
	for i := 0; i < wordLength; i++ {
		index := w[i] - 'a'
		if temp.children[index] == nil {
			return false
		}
		temp = temp.children[index]
	}
	if temp.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := Trie{}
	myTrie.root = &trieNode{}

	myTrie.buildTrie("sanika")
	myTrie.buildTrie("sarika")
	myTrie.buildTrie("anamika")
	myTrie.buildTrie("saranya")
	myTrie.buildTrie("abi")
	myTrie.buildTrie("lavanya")
	myTrie.root.PrintTrie(" ")
	fmt.Println(myTrie.Search("lavanya"))
	myTrie.searchByPrefix(myTrie.root, "s")

}
