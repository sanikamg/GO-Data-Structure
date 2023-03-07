package main

import (
	"fmt"
)

func reverseArray(arr [7]int) []int {
	n := len(arr) - 1
	reversed := make([]int, n)
	for i := 0; i < n; i++ {
		reversed[i] = arr[n-i-1]
	}
	return reversed
}

func sumTen(array [7]int) (int, int) {

	for i := 0; i < 7; i++ {
		for j := i + 1; j < 7; j++ {
			if array[i]+array[j] == 10 {
				return array[i], array[j]

			}
		}
	}
	return 0, 0

}

func main() {
	array := [7]int{9, 5, 4, 3, 6, 8}

	i, j := sumTen(array)
	fmt.Printf("the numbers are %d , %d\n", i, j)
	reversed := reverseArray(array)
	fmt.Println(reversed)
}
