package main

import "fmt"

func buildMinHeap(array []int) {
	lastNonLeaf := len(array)/2 - 1
	for i := lastNonLeaf; i >= 0; i-- {
		minHeapify(array, i)
	}
}
func minHeapify(array []int, i int) {
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
		minHeapify(array, min)
	}
}

func minHeapSort(array []int) {
	length := len(array);
	for i := length - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		minHeapify(array[:i], 0)
	}

}

func main() {
	array := []int{100, 9999, 78, 999, 36}
	fmt.Println(array)
	buildMinHeap(array)
	fmt.Println(array)
	minHeapSort(array)
	fmt.Println(array)
}
