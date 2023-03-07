package main

import "fmt"

func maxHeapify(array []int, i int) {
	length := len(array)
	largest := i
	left := (2 * i) + 1
	right := (2 * i) + 2
	if left < length && array[left] > array[largest] {
		largest = left
	}
	if right < length && array[right] > array[largest] {
		largest = right
	}
	if i != largest {
		array[i], array[largest] = array[largest], array[i]
		maxHeapify(array, largest)
	}
}

func buildMaxHeap(array []int) {
	length := len(array)
	lastNonLeaf := (length/2 - 1)
	for i := lastNonLeaf; i >= 0; i-- {
		maxHeapify(array, i)
	}
}

func heapSort(array []int) []int {
	var result []int
	length := len(array) - 1
	for i := length; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		result = append(result, array[i])
		maxHeapify(array[:i], 0)
	}
	return result
}

func delete(array *[]int) {
	length := len(*array) - 1
	(*array)[0] = (*array)[length]
	*array = (*array)[:length-1]
	maxHeapify(*array, 0)
}

func main() {
	array := []int{10, 90, 2022, 206, -22, -55, 777, 99999, 0, 999999999}
	fmt.Println(array)
	buildMaxHeap(array)
	fmt.Println(array)
	SortedHeap := heapSort(array)
	fmt.Println(SortedHeap)
	delete(&array)
	fmt.Println(array)
}
