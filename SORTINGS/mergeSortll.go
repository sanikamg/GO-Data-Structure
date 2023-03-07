package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// SplitList splits the linked list into two halves
func SplitList(head *Node) (*Node, *Node) {
	slow := head
	fast := head.Next

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
	}

	mid := slow.Next
	slow.Next = nil

	return head, mid
}

// MergeSort sorts a linked list using the Merge Sort algorithm
func MergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	first, second := SplitList(head)
	first = MergeSort(first)
	second = MergeSort(second)

	return Merge(first, second)
}

// Merge merges two sorted linked lists into one
func Merge(first, second *Node) *Node {
	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	var result *Node

	if first.Value <= second.Value {
		result = first
		result.Next = Merge(first.Next, second)
	} else {
		result = second
		result.Next = Merge(first, second.Next)
	}

	return result
}

func main() {
	head := &Node{Value: 5, Next: &Node{Value: 20, Next: &Node{Value: 4, Next: &Node{Value: 3, Next: &Node{Value: 30}}}}}
	head = MergeSort(head)

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}
