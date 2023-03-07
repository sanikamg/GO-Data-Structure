package main

import "fmt"

func main() {
	var num int
	var temp int
	array1 := [7]int{1, 9, 6, 6, 7, 8, 6}
	fmt.Println("Enter the number you want to sort : ")
	fmt.Scan(&num)

	j := len(array1)
	fmt.Println(j)
	for i, a := range array1 {
		if j == i {
			break
		}
		if a == num {
			if array1[j-1] == num {
				j--

			}

			fmt.Println(j)
			if array1[j-1] != num {
				temp = array1[i]
				array1[i] = array1[j-1]
				array1[j-1] = temp
				j--
			}
		}

	}
	fmt.Println(array1)
}

// temp = a
// 			a = array1[j-1]
// 			array1[j-1] = temp
// 			if array1[j-2] == num {
// 				j = j - 1
// 			} else {
// 				j--
// 			}
