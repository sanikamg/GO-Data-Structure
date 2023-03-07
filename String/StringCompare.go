package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	var str2 string

	fmt.Println("Enter any string : ")
	fmt.Scan(&str1)

	fmt.Println("Enter any string : ")
	fmt.Scan(&str2)

	num := strings.Compare(str1, str2)

	if num == 0 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}

}
