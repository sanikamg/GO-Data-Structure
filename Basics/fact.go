package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Enter your number: ")
	fmt.Scanf("%d", &n)
	factorial := fact(n)
	fmt.Printf("factorial is %d", factorial)

}

func fact(n int) int {
	var res int
	if n == 0 {
		res = 1
	} else {
		res = n * fact(n-1)
	}
	return res
}
