package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	var delete int

	fmt.Println("Enter your array size : ")
	fmt.Scan(&size)

	var array = make([]int, size)
	fmt.Println("Enter your array values : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])

	}

	fmt.Println("Enter the values you want to delete : ")
	fmt.Scan(&delete)
	i := sort.SearchInts(array, delete)

	array = append(array[:i], array[i+1:]...)

	fmt.Println(array)

}
