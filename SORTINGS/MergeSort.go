package main

import "fmt"

func mergeSort(array []int, lb int, ub int) {
	if lb < ub {
		mid := (lb + ub) / 2
		mergeSort(array, lb, mid)
		mergeSort(array, mid+1, ub)
		merge(array, lb, mid, ub)
	}
}

func merge(array []int, lb int, mid int, ub int) {
	i := lb
	j := mid + 1
	k := lb
	tempArray := make([]int, ub+1)

	for i <= mid && j <= ub {
		if array[i] < array[j] {
			tempArray[k] = array[i]
			i++
		} else {
			tempArray[k] = array[j]
			j++
		}
		k++
	}
	if i > mid {
		for j <= ub {
			tempArray[k] = array[j]
			j++
			k++
		}
	} else {
		for i <= mid {
			tempArray[k] = array[i]
			i++
			k++
		}
	}
	for i := lb; i <= ub; i++ {
		array[i] = tempArray[i]
	}
}

func main() {
	array := []int{10, 20, 100, 90, 50, 30, 60, 70, 80, 40}
	fmt.Println(array)
	lb := 0
	ub := len(array) - 1
	mergeSort(array, lb, ub)
	fmt.Println(array)

}
