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

func BubbleSort(array []int, size int) {
	num := size - 1
	// forloop for select next start
	for i := 0; i < size; i++ {
		// forloop for swapping and the smallest element come to first
		for j := 0; j < num; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		num--

	}
	fmt.Println(array)

}

func main() {
	array, size := arrayCreation()
	BubbleSort(array, size)
}
