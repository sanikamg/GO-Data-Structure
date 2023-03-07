package main

import "fmt"

func Encryption(str string) {
	for _, s := range str {
		char := s + 25
		if char < 123 {
			fmt.Printf("%s", string(char))
		} else {
			char = char - 26
			fmt.Printf("%s", string(char))
		}

	}
}

func main() {

	var str string

	fmt.Println("Enter any string : ")
	fmt.Scan(&str)
	Encryption(str)

}
