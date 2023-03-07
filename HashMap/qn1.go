//Given a list of words, write a function to determine if a given word is in the list.

package main

import "fmt"

func contains(list []string, str string) bool {
	hashTable := make(map[string]bool)

	for _, w := range list {
		hashTable[w] = true
	}
	_, exists := hashTable[str]
	return exists
}

func main() {
	var list = []string{"hi", "hello", "sanika"}
	str := "sanika"
	result := contains(list, str)
	fmt.Println(result)
}
