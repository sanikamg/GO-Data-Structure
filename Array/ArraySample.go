package main

import "fmt"

func main() {
	//init
	array := [4]int{1, 2, 3, 4}
	// set-space complexity O(1),time Complexity O(1)
	fmt.Println(array[0]) //get-space complexity O(1),time complexity O(1)
	array[0] = 5
	fmt.Println(array[0]) //get
	//traversal,time-O(n),space-O(1)
	for i := 0; i <= 3; i++ {
		fmt.Println(array[i])
	}
	//insert at the end time-O(n),space-O(1)

	//Dynamic array-we can go with dynamic size

}
