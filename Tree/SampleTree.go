package main

import "fmt"

type tree struct {
	root *treeNode
}
type treeNode struct {
	key   int
	left  *treeNode
	right *treeNode
}

func createNode(key int) *treeNode {
	newNode := new(treeNode)
	newNode.key = key
	newNode.left = nil
	newNode.right = nil
	return newNode
}
func (node *treeNode) printPreorder() {
	if node != nil {
		fmt.Printf(" %d ", node.key)
		node.left.printPreorder()
		node.right.printPreorder()
	}
	return
}

func treeIdentical(root1, root2 *treeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.key != root2.key {
		return false
	}

	return treeIdentical(root1.left, root2.left) && treeIdentical(root1.right, root2.right)
}

func isBSt(root *treeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.key <= min || root.key >= max {
		return false
	}
	return isBSt(root.left, min, root.key) && isBSt(root.right, root.key, max)
}

func main() {

	root1 := createNode(10)
	root1.left = createNode(5)
	root1.right = createNode(40)
	root1.right.left = createNode(35)
	root1.right.right = createNode(75)
	root1.printPreorder()
	fmt.Println(" ")
	root2 := createNode(110)
	root2.left = createNode(220)
	root2.right = createNode(40)
	root2.right.left = createNode(45)
	root2.right.right = createNode(75)
	root2.printPreorder()
	fmt.Println(treeIdentical(root1, root2))
	if isBSt(root1, -1000, 1000) {
		fmt.Println("valid BST")
	} else {
		fmt.Println("invalid BST")
	}
}
