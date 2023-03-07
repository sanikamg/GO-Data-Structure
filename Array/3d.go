package main

import "fmt"

func main() {
	// Create a 3D array
	arr := [2][3][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{13, 14, 15, 16},
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}

	// Iterate over the array and print each element
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			for k := 0; k < len(arr[i][j]); k++ {
				fmt.Printf("%d ", arr[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
