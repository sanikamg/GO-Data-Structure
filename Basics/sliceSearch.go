package main

import (
	"fmt"
	"sort"
)

// Main function
func main() {

	// Creating and initializing
	// slice of ints using the
	// shorthand declaration
	slice_1 := []int{34, 67, 78, 10, 43, 67, 8}
	slice_2 := []int{100, 500, 300, 600, 900, 1000}
	fmt.Println(len(slice_1))
	var f1, f2, f3 int
	f1 = 67
	f2 = 300
	f3 = 100

	// Sorting the given slice of ints
	sort.Ints(slice_1)
	sort.Ints(slice_2)

	// Displaying the slices
	fmt.Println("Slice 1: ", slice_1)
	fmt.Println("Slice 2: ", slice_2)

	// Searching a int type element
	// in the given slice using
	// the SearchInts function
	res1 := sort.SearchInts(slice_1, f1)
	res2 := sort.SearchInts(slice_1, f2)
	res3 := sort.SearchInts(slice_2, f3)

	// Displaying the results
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)

}
