package main

import "fmt"

func arrayCreation() []int {
	var size int
	fmt.Println("Ente your array size : ")
	fmt.Scan(&size)

	var array = make([]int, size)
	fmt.Println("Enter your array values : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(array)
	return array
}

func quickSort(array []int, lb int, ub int) {
	if lb < ub {
		pivot := partition(array, lb, ub)
		//recursively calling for sorting elements less than pivot element
		quickSort(array, lb, pivot-1)
		//recursively calling for sorting elements greater than pivot element
		quickSort(array, pivot+1, ub)
	}
}

func partition(array []int, lb int, ub int) int {
	pivot := array[lb]
	start := lb
	end := ub

	for start < end {
		// for increment start until element is greater than pivot element
		for array[start] <= pivot && start < ub {
			start++
		}
		// for decrement end until element is less than pivot
		for array[end] > pivot && end > lb {
			end--
		}
		// for swapping if element is unordered position
		if start < end {
			array[start], array[end] = array[end], array[start]
		}
	}
	array[lb], array[end] = array[end], pivot
	//return the pivot element correct position
	return end
}

func main() {
	array := arrayCreation()
	lb := 0
	ub := len(array) - 1
	quickSort(array, lb, ub)
	fmt.Println(array)
}
