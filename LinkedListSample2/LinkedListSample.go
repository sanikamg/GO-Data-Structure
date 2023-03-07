package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
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

		if temp.next == nil {
			break
		} else {
			temp = temp.next.next
		}

	}
}

func (ll *LinkedList) Delete(num int) {
	temp := ll.head
	for temp != nil {
		flag := temp
		temp = temp.next
		if temp.data == num {
			flag.next = temp.next
			return
		}
	}
}

func main() {

	list := LinkedList{}
	list.CreateList(10)
	list.CreateList(20)
	list.CreateList(30)
	list.CreateList(40)
	list.CreateList(50)

	list.Display()

}
