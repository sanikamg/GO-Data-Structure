package main

import "fmt"

func arrayCreation() ([]int, int) {
	var size int
	fmt.Println("Ente your array size : ")
	fmt.Scan(&size)

	var array = make([]int, size)
	fmt.Println("Enter your array values : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(array)
	return array, size
}

func InsertionSort(array []int, size int) {
	// for insert next element
	for index, _ := range array {
		if index == 0 {
			continue
		} else {
			// for check the previous elements
			for i := index; i > 0; i-- {
				if array[i] < array[i-1] {
					temp := array[i]
					array[i] = array[i-1]
					array[i-1] = temp
				}
			}

		}
	}
	fmt.Println(array)

}

func main() {
	array, size := arrayCreation()
	InsertionSort(array, size)
}
