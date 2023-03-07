package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func ArrayCreation() ([]int, int) {
	var size int
	fmt.Println("Enter your array size : ")
	fmt.Scan(&size)
	var array = make([]int, size)
	fmt.Println("Enter ", size, "array values : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(array)
	return array, size
}

func (ll *LinkedList) Insert(array []int, size int) {

	for i := 0; i < size; i++ {
		newNode := new(Node)
		newNode.data = array[i]
		if ll.head == nil {
			ll.head = newNode
		} else {
			temp := ll.head
			for temp != nil {
				if temp.next == nil {
					temp.next = newNode
					break
				} else {
					temp = temp.next
				}
			}
		}
	}
}

func (ll *LinkedList) Display() {
	temp := ll.head
	for temp != nil {
		fmt.Printf("%d-->", temp.data)
		temp = temp.next
	}
}

func main() {
	var choice int
	list := LinkedList{}
	for true {
		fmt.Println("\nSelect any operation : ")

		fmt.Println("1. insertion")
		fmt.Println("2.Display")
		fmt.Println("3.Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:

			array, size := ArrayCreation()
			list.Insert(array, size)
		case 2:
			list.Display()

		case 3:
			os.Exit(0)

		}

	}
}
