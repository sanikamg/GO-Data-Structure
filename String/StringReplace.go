// Write a function to replace each alphabet in the given string with another alphabet occurring at the n-th position from each of them.
package main

import "fmt"

func main() {
	var str string

	fmt.Println("Enter any string : ")
	fmt.Scan(&str)
	num := 3

	//num := 2
	for i := 0; i < len(str); i++ {
		if i+num >= len(str) {
			break
		} else {
			fmt.Printf("%s", string(str[i+num]))
		}

	}
}
