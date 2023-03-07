package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	prev *Node
	next *Node
}
type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) CreateList(num int) {
	newNode := new(Node)
	newNode.data = num

	if ll.head == nil {
		ll.head = newNode
	} else {
		temp := ll.head
		for temp != nil {
			if temp.next == nil {
				temp.next = newNode
				newNode.prev = temp
				ll.tail = newNode
				return
			}
			temp = temp.next

		}

	}
}

func (ll *LinkedList) Display() {

	temp := ll.head

	for temp != nil {
		fmt.Printf("%d -->", temp.data)
		temp = temp.next

	}

}
func (ll *LinkedList) Reverse() {
	temp := ll.tail
	for temp != nil {
		fmt.Printf("%d -->", temp.data)
		temp = temp.prev

	}

}
func main() {
	var choice int
	list := LinkedList{}

	for true {
		fmt.Println("\nEnter your choice : ")
		fmt.Println("1.Insert value")
		fmt.Println("2.Display")
		fmt.Println("3.Reverse")
		fmt.Println("4.Exit")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter your value : ")
			fmt.Scan(&num)
			list.CreateList(num)
		case 2:
			list.Display()
		case 3:
			list.Reverse()

		default:
			os.Exit(0)
		}
	}

}
