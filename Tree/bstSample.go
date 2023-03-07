package main

import "fmt"

//tree node structure
type bstTreeNode struct {
	data  int
	left  *bstTreeNode
	right *bstTreeNode
}
type BstTree struct {
	root *bstTreeNode
}

func (bst *BstTree) insert(num int) {
	newNode := new(bstTreeNode)
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

func (bst *BstTree) delete(num int) {
	var parent *bstTreeNode
	temp := bst.root
	if temp != nil && temp.data != num {
		parent = temp
		if temp.data > num {
			temp = temp.left
		} else {
			temp = temp.right
		}

	}

	if temp == nil {
		fmt.Println("element not found")
	}

	if temp.right == nil && temp.left == nil {
		if parent == nil {
			bst.root = nil
		} else if parent.left == temp {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if temp.right == nil && temp.left != nil {
		if parent == nil {
			bst.root = nil
		} else if parent.left == temp {
			parent.left = temp.left
		} else {
			parent.right = temp.left
		}
	} else if temp.right != nil && temp.left == nil {
		if parent == nil {
			bst.root = nil
		} else if parent.left == temp {
			parent.left = temp.right
		} else {
			parent.right = temp.right
		}
	} else {
		inOrderSuccessor := temp.right
		for inOrderSuccessor.left != nil {
			inOrderSuccessor = inOrderSuccessor.left
		}
		bst.delete(inOrderSuccessor.data)
		temp.data = inOrderSuccessor.data
	}

}

func (bst *BstTree) printPreOrder(node *bstTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf(" %d-->", node.data)
	bst.printPreOrder(node.left)
	bst.printPreOrder(node.right)
}

func (bst BstTree) printInorder(node *bstTreeNode) {
	if node == nil {
		return
	}
	bst.printInorder(node.left)
	fmt.Printf("%d-->", node.data)
	bst.printInorder(node.right)
}

func (bst *BstTree) sumofNode(node *bstTreeNode) int {
	if node == nil {
		return 0
	}
	return node.data + bst.sumofNode(node.left) + bst.sumofNode(node.right)
}

func main() {
	BST := BstTree{}
	BST.insert(10)
	BST.insert(20)
	BST.insert(80)
	BST.insert(5)
	BST.insert(90)
	BST.insert(1)
	BST.printPreOrder(BST.root)
	fmt.Println(" ")
	BST.printInorder(BST.root)
	sum := BST.sumofNode(BST.root)
	fmt.Printf("\nsum of these nodes are : %v", sum)
	BST.delete(10)
	BST.printPreOrder(BST.root)

}
