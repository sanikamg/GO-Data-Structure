package main

import "fmt"

func main() {
	// create
	nameAge := map[string]int{
		"sanika":  25,
		"sandhya": 27,
		"vishnu":  27,
	}
	//insert
	nameAge["latha"] = 45
	nameAge["gopalan"] = 52

	//delete
	delete(nameAge, "latha")
	//search
	_, exists := nameAge["latha"]
	if exists {
		fmt.Println("exists")
	} else {
		fmt.Println("not exists")
	}
	fmt.Println(nameAge)
}
