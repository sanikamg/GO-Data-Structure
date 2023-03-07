package main

import (
	"fmt"
	"os"
)

type LinkedList struct {
	number int
	next   *LinkedList
}

func (node *LinkedList) insert(num int) {
	var temp = new(LinkedList)
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp

	} else {
		var p = new(LinkedList)
		p = node
		for p.next != nil {
			p = p.next

		}
		p.next = temp

	}

}

func (node *LinkedList) Display() {
	var p = new(LinkedList)
	p = node.next
	for p != nil {
		fmt.Printf("%d -->", p.number)
		p = p.next
	}
}

func (node *LinkedList) insertBegin(num int) {

}

func main() {
	head := new(LinkedList)
	var choice int
	for true {
		fmt.Println("\nEnter your choice : ")
		fmt.Println("1.insert value")
		fmt.Println("2:inser begin of the list")
		fmt.Println("3.Display value")
		fmt.Println("4.Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var data int
			fmt.Println("Enter your value : ")
			fmt.Scan(&data)
			head.insert(data)

		case 2:
			var num int
			fmt.Println("Enter your value : ")
			fmt.Scan(&num)
			head.insertBegin(num)
		case 3:
			head.Display()

		default:
			os.Exit(0)

		}
	}
}
