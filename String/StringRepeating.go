package main

import "fmt"

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func CountAlphabet(str string) {
	var slice []string

	for _, s := range str {
		if contains(slice, string(s)) == true {
			continue
		} else {
			count := 0
			for _, value := range str {
				if s == value {
					count++
				}
			}
			fmt.Printf("%d%s ", count, string(s))
			slice = append(slice, string(s))
		}

	}
}

func main() {

	var str string

	fmt.Println("Enter any string : ")
	fmt.Scan(&str)
	CountAlphabet(str)

}
