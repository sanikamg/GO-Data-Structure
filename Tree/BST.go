package main

import "fmt"

type bstNode struct {
	data  int
	left  *bstNode
	right *bstNode
}

type bstTree struct {
	root *bstNode
}

func (bst *bstTree) insert(num int) {
	newNode := new(bstNode)
	newNode.data = num
	if bst.root == nil {
		bst.root = newNode
	} else {
		temp := bst.root
		for temp != nil {
			if temp.data > num {
				if temp.left == nil {
					temp.left = newNode
					return
				} else {
					temp = temp.left
				}
			} else {
				if temp.right == nil {
					temp.right = newNode
					return
				} else {
					temp = temp.right
				}
			}
		}
	}
}

func validateBST(root *bstNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.data <= min || root.data >= max {
		return false
	}

	return validateBST(root.left, min, root.data) && validateBST(root.right, root.data, max)
}
func (bst *bstTree) printPreOrder(node *bstNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d-->", node.data)
	bst.printPreOrder(node.left)
	bst.printPreOrder(node.right)
}
func main() {

	BST := bstTree{}
	BST.insert(10)
	BST.insert(30)
	BST.insert(20)
	BST.insert(15)
	BST.insert(25)
	BST.insert(5)
	BST.insert(40)
	BST.insert(35)
	BST.insert(50)
	BST.insert(45)
	BST.insert(60)
	BST.printPreOrder(BST.root)
	if validateBST(BST.root, -1000, 1000) {
		fmt.Println("Valid BST")
	} else {
		fmt.Println("Invalid BST")
	}

}
