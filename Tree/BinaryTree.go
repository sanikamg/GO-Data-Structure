package main

import "fmt"

type btNode struct {
	data  int
	left  *btNode
	right *btNode
}

func createNode(num int) *btNode {
	newNode := new(btNode)
	newNode.data = num
	newNode.left = nil
	newNode.right = nil
	return newNode
}

func (node *btNode) printPreorder() {
	if node != nil {
		fmt.Printf("%d-->", node.data)
		node.left.printPreorder()
		node.right.printPreorder()
	}
}

func (node *btNode) sumOfNodes() int {
	if node == nil {
		return 0
	}
	return node.data + node.left.sumOfNodes() + node.right.sumOfNodes()
}

func validateBST(root *btNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.data <= min || root.data >= max {
		return false
	}

	return validateBST(root.left, min, root.data) && validateBST(root.right, root.data, max)
}

func main() {
	root := createNode(1)
	root.left = createNode(2)
	root.right = createNode(3)
	root.left.left = createNode(4)
	root.left.right = createNode(5)
	root.right.left = createNode(6)
	root.right.right = createNode(7)
	root.printPreorder()
	sum := root.sumOfNodes()

	fmt.Println(sum)
	if validateBST(root, -1000, 1000) {
		fmt.Println("Valid BST")
	} else {
		fmt.Println("Invalid BST")
	}
}
