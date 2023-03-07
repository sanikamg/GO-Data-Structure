package main

import (
	"fmt"
	"os"
)

func createHeap(array *[]int, value int) {
	*array = append(*array, value)
	i := len(*array) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if (*array)[parent] < value {
			(*array)[parent], (*array)[i] = (*array)[i], (*array)[parent]
			i = parent
		} else {
			break
		}
	}
}

func main() {
	array := make([]int, 0)
	var value, choice int
	for {
		fmt.Println("Enter any option : ")
		fmt.Println("1 : insert")
		fmt.Println("2 : display")
		fmt.Println("3 : Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter the value you want to insert : ")
			fmt.Scan(&value)
			createHeap(&array, value)

		case 2:
			fmt.Println(array)
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}

	}
}
