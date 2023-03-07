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

func (ll *LinkedList) insertNode(num int) {

	newNode := new(Node)
	newNode.data = num

	if ll.head == nil {
		ll.head = newNode
		return
	}
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

func (ll *LinkedList) Display() {

	temp := ll.head

	for temp != nil {
		fmt.Printf("%d -->", temp.data)
		temp = temp.next

	}

}

func (ll *LinkedList) insertBegin(num int) {
	newNode := new(Node)
	newNode.data = num
	temp := ll.head
	ll.head = newNode
	newNode.next = temp

}

func (ll *LinkedList) insertMiddle(num, num1 int) {

	newNode := new(Node)
	newNode.data = num1
	temp := ll.head
	for temp != nil {
		prev := temp
		if temp.data == num {
			flag := prev.next
			prev.next = newNode
			newNode.next = flag
			return
		}
		temp = temp.next

	}

}

func (ll *LinkedList) DeleteNode(num int) {
	temp := ll.head
	for temp != nil {
		if temp.next.data == num {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}
func (ll *LinkedList) SearchNode(num int) {
	temp := ll.head
	for temp != nil {
		fmt.Println(temp)
		if temp.data == num {
			fmt.Println("Element found.")
			return
		}

		temp = temp.next

	}
	fmt.Println("Element not found.....")
}

func (ll *LinkedList) ReplaceNode(num, num1 int) {
	newNode := new(Node)
	newNode.data = num1
	temp := ll.head
	for temp != nil {
		if temp.next.data == num {
			temp.next.data = num1
			return
		}
		temp = temp.next
	}
}

func (ll *LinkedList) sortList() {

	temp := ll.head

	for temp != nil {
		num1 := temp.next
		for num1 != nil {
			if num1.data < temp.data {
				flag := temp.data
				temp.data = num1.data
				num1.data = flag
			}
			num1 = num1.next
		}
		fmt.Printf("%d-->", temp.data)
		temp = temp.next
	}

}

func (ll *LinkedList) DuplicateDelete() {
	temp := ll.head

	for temp != nil {
		num1 := temp.next
		for num1 != nil {
			if num1.data == temp.data {
				ll.DeleteNode(num1.data)
			}
			num1 = num1.next
		}

		temp = temp.next
	}
}

func main() {
	var choice int
	list := LinkedList{}

	for true {
		fmt.Println("\nEnter your choice : ")
		fmt.Println("1.insert value")
		fmt.Println("2.insert begin of the list")
		fmt.Println("3.insert middle of the list")
		fmt.Println("4.Delete value")
		fmt.Println("5.Display value")
		fmt.Println("6.search value")
		fmt.Println("7.Replace value")
		fmt.Println("8.SortList")
		fmt.Println("9.Delete duplicate node")
		fmt.Println("10.Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter your value : ")
			fmt.Scan(&num)
			list.insertNode(num)

		case 2:
			var num int
			fmt.Println("Enter your value : ")
			fmt.Scan(&num)
			list.insertBegin(num)
		case 3:
			var num int
			var num1 int
			fmt.Println("Enter your number after that you want to insert : ")
			fmt.Scan(&num)
			fmt.Println("Enter your number that want to insert : ")
			fmt.Scan(&num1)
			list.insertMiddle(num, num1)
		case 4:
			var num int
			fmt.Println("Enter your value to delete : ")
			fmt.Scan(&num)
			list.DeleteNode(num)
		case 5:
			list.Display()
		case 6:
			var num int
			fmt.Println("Enter the value you want to search : ")
			fmt.Scan(&num)
			list.SearchNode(num)
		case 7:

			var num int
			var num1 int
			fmt.Println("Enter your number  that you want to replace : ")
			fmt.Scan(&num)
			fmt.Println("Enter your number that want to add : ")
			fmt.Scan(&num1)
			list.ReplaceNode(num, num1)

		case 8:
			list.sortList()
		case 9:
			list.DuplicateDelete()

		default:
			os.Exit(0)

		}
	}

}
