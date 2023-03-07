package main

import "fmt"

func Fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}

func main() {
	var num int
	fmt.Println("Enter any number : ")
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fibonacci := Fibonacci(i)
		fmt.Printf("%d ", fibonacci)
	}

}
