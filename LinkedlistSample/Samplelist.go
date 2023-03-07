package main

import "fmt"

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

func (ll *LinkedList) display() {
	temp := ll.head
	for temp != nil {
		fmt.Printf("%d-->", temp.data)
		temp = temp.next
	}
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.display()
}
