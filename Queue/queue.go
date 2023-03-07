package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}
type queue struct {
	top *Node
}

func (q *queue) enqueue(num int) {
	newnode := new(Node)
	newnode.data = num
	if q.top == nil {
		q.top = newnode
	} else {
		temp := q.top
		for temp != nil {
			if temp.next == nil {
				temp.next = newnode
				return
			} else {
				temp = temp.next
			}
		}
	}
}

func (q *queue) dequeue() {
	if q.top == nil {
		fmt.Println("Queue is in overflow")
	} else {
		q.top = q.top.next
	}

}

func (q *queue) display() {
	temp := q.top
	for temp != nil {
		fmt.Printf("[%d]\n", temp.data)
		temp = temp.next
	}
}

func main() {

	list := queue{}
	var choice int
	for true {
		fmt.Println("\nselect your choice : ")
		fmt.Println(" 1:enqueue  ")
		fmt.Println(" 2:dequeue ")
		fmt.Println(" 3:display ")
		fmt.Println(" 4:Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter the Number you want to enqueue : ")
			fmt.Scan(&num)
			list.enqueue(num)
		case 2:
			list.dequeue()
		case 3:
			list.display()

		case 4:
			os.Exit(0)

		}
	}

}
