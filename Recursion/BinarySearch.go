package main

import "fmt"

func ArrayCreation() ([]int, int) {
	var size int
	fmt.Println("Enter your array size : ")
	fmt.Scan(&size)
	var array = make([]int, size)
	fmt.Println("Enter ", size, "array values : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(array)
	return array, size
}

func ArraySort(array []int, size int) []int {
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if array[i] > array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}

		}
	}
	fmt.Println(array)
	return array
}

func BinarySearch(array []int, low int, high int) int {
	var num int
	fmt.Println("Enter the element you want to search : ")
	fmt.Scan(&num)
	middle := (low + high) / 2

	if array[middle] == num {
		fmt.Println("Element found in ", middle, " position.")
		return 1
	} else if num > array[middle] {
		return BinarySearch(array, middle+1, high)
	} else if num < array[middle] {
		return BinarySearch(array, low, middle-1)
	} else {
		fmt.Println("Element not found : ")
	}
	return 1
}

func main() {
	array1, size := ArrayCreation()
	low := 0
	high := size
	array := ArraySort(array1, size)
	BinarySearch(array, low, high)
}
