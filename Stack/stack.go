package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}
type Stack struct {
	top *Node
}

func (s *Stack) push(num int) {
	newNode := new(Node)
	newNode.data = num
	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (s *Stack) pop() {
	if s.top == nil {
		fmt.Println("Stack underflow")
	} else {
		s.top = s.top.next
	}
}

func (s *Stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Printf("[%d]\n", temp.data)
		temp = temp.next
	}
}

func main() {
	list := Stack{}
	var choice int
	for true {
		fmt.Println("\nselect your choice : 1:Push  2:Pop 3:Exit")
		fmt.Println(" 1:Push  ")
		fmt.Println(" 2:Pop ")
		fmt.Println(" 3:display ")
		fmt.Println(" 4:Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter the Number you want to push : ")
			fmt.Scan(&num)
			list.push(num)
		case 2:
			list.pop()
		case 3:
			list.display()

		case 4:
			os.Exit(0)

		}
	}

}
