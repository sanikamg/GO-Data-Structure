package main

import "fmt"

func main() {
	var str string

	fmt.Println("Enter any string : ")
	fmt.Scan(&str)
	fmt.Println(len(str))
	num := len(str)
	var slice []string
	for i := num - 1; i >= 0; i-- {
		slice = append(slice, string(str[i]))

	}
	fmt.Println(slice)

}
