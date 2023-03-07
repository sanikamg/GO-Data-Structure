package main

import "fmt"

func buildMaxHeap(array []int) {
	length := len(array)
	lastNonLeaf := (length / 2) - 1
	for i := lastNonLeaf; i >= 0; i-- {
		maxHeapify(array, i)
	}
}

func maxHeapify(array []int, i int) {
	length := len(array)
	min := i
	left := (2 * i) + 1
	right := (2 * i) + 2
	if left < length && array[left] < array[min] {
		min = left
	}
	if right < length && array[right] < array[min] {
		min = right
	}
	if i != min {
		array[i], array[min] = array[min], array[i]
		maxHeapify(array, min)
	}

}

func main() {
	array := []int{1, 20, 3, 400, 55, 6}
	buildMaxHeap(array)
	fmt.Println(array)

}
