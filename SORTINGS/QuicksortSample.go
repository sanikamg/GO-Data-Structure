package main

import "fmt"

func quickSort(array [5]int, lb int, ub int) {
	if lb < ub {
		pivotLoc := partition(array, lb, ub)
		quickSort(array, lb, pivotLoc-1)
		quickSort(array, pivotLoc+1, ub)
	}

}

func partition(array [5]int, lb int, ub int) int {
	pivot := array[lb]
	start := lb
	end := ub

	for start < end {

		for array[start] <= pivot && start < ub {
			start++
		}
		for array[end] > pivot && start > lb {
			end--
		}
		if start < end {
			array[end], array[start] = array[start], array[end]
		}

	}

	array[lb], array[end] = array[end], pivot
	return end
}

func main() {

	var array [5]int

	fmt.Println("Enter your array elements : ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	lb := 0
	ub := 4
	quickSort(array, lb, ub)
	fmt.Println(array)

}
