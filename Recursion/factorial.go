package main

import "fmt"

func Fact(num int) int {
	if num < 1 {
		return 1
	}
	return num * Fact(num-1)
}

func main() {
	var num int
	fmt.Println("Enter any number : ")
	fmt.Scan(&num)
	factorial := Fact(num)
	fmt.Println(factorial)

}
