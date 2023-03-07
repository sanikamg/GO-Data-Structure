package main

import "fmt"

func Search(array [5]int, num int) {
	for index, value := range array {
		if num == value {
			fmt.Printf("Element is found at %d position", index)
			return
		}

	}
	fmt.Println("This element not found in array")

}

func main() {
	var num int
	var array = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Enter the element you want to search : ")
	fmt.Scan(&num)
	Search(array, num)

}
