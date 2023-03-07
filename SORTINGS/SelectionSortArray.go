package main

import "fmt"

var num int

var i int

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

func SelectionSort(array []int) {

	for k := 0; k < len(array)-1; k++ {
		small := k

		for i := k + 1; i < len(array); i++ {
			if array[i] < array[small] {
				small = i
			}

		}

		//fmt.Println(small)

		temp := array[k]
		array[k] = array[small]
		array[small] = temp

	}

	fmt.Println(array)

}

func main() {
	array := arrayCreation()
	SelectionSort(array)
}
