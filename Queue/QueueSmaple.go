package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type queue struct {
	front *Node
	rear  *Node
}

func (q *queue) enqueue(num int) {
	newNode := new(Node)
	newNode.data = num
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *queue) dequeue() {
	if q.front == nil {
		fmt.Println("queue is in underflow")
	} else {
		q.front = q.front.next
	}
}

func (q *queue) display() {
	temp := q.front
	for temp != nil {
		fmt.Printf("\n[%d]", temp.data)
		temp = temp.next
	}
}

func main() {
	Queue := queue{}
	Queue.enqueue(1)
	Queue.enqueue(2)
	Queue.display()
	Queue.dequeue()
	fmt.Println("\n-----------")
	fmt.Println("Queue after enqueue..")
	Queue.display()
}
