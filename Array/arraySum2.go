package main

import (
	"fmt"
	"sort"
)

func sum(array [7]int) (int, int) {
	var slice []int
	target := 10
	for _, a := range array {

		res := sort.SearchInts(slice, a)
		num := target - a

		if len(slice) == res {
			slice = append(slice, a)
		}

		match := sort.SearchInts(slice, num)
		if match != len(slice) {
			return num, a

		}

	}
	return 0, 0
}
func main() {

	array := [7]int{7, 4, 2, 6, 3, 8, 0}
	num1, num2 := sum(array)
	fmt.Println("numbers are : ", num1, num2)

}

// if len(slice) == 0 {
// 	slice = append(slice, 10-a)
// 	fmt.Println(len(slice))
// } else if res == len(slice) {
// 	slice = append(slice, 10-a)
// 	fmt.Println(slice)
// } else {
// 	if 10-a == slice[res] {
// 		fmt.Println(10-a, slice[res])
// 		fmt.Println(slice)
// 	}
// }
