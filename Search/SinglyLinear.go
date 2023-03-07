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

func (ll *LinkedList) Insert(num int) {
	newNode := new(Node)
	newNode.data = num
	if ll.head == nil {
		ll.head = newNode
	} else {
		temp := ll.head
		for temp != nil {
			if temp.next == nil {
				temp.next = newNode
				return
			} else {
				temp = temp.next
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

func (ll *LinkedList) Search(num int) {
	temp := ll.head
	for temp != nil {
		if temp.data == num {
			fmt.Println("Element is found")
			return
		}
		temp = temp.next
	}
	fmt.Println("Element not found")

}

func main() {
	var choice int
	list := LinkedList{}
	for true {
		fmt.Println("\nSelect any operation : ")

		fmt.Println("1. insertion")
		fmt.Println("2.Display")
		fmt.Println("3.Linear Search")
		fmt.Println("4.Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var num int
			fmt.Println("Enter the number : ")
			fmt.Scan(&num)
			list.Insert(num)
		case 2:
			list.Display()
		case 3:
			var num int
			fmt.Println("Enter the number you want to search : ")
			fmt.Scan(&num)
			list.Search(num)
		case 4:
			os.Exit(0)

		}

	}
}
