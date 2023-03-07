package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bstTree struct {
	root *Node
}

func (bst *bstTree) insert(num int) {
	newNode := new(Node)
	newNode.data = num
	if bst.root == nil {
		bst.root = newNode
	} else {
		temp := bst.root
		for temp != nil {
			if newNode.data < temp.data {
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

func (bst *bstTree) contains(num int) {
	temp := bst.root
	for temp != nil {
		if temp.data == num {
			fmt.Println("Element found")
			return
		} else if temp.data > num {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	fmt.Println("Element not found")
}

func (bst *bstTree) delete(num int) {
	var parent *Node
	temp := bst.root
	for temp != nil && temp.data != num {
		parent = temp
		if temp.data > num {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	if temp == nil {
		fmt.Println("Element not found")
		return
	}

	if temp.left == nil && temp.right == nil {
		if parent == nil {
			bst.root = nil
		} else if parent.left == temp {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if temp.left != nil && temp.right == nil {
		if parent == nil {
			bst.root = temp
		} else if parent.left == temp {
			parent.left = temp.left
		} else {
			parent.right = temp.left
		}
	} else if temp.right != nil && temp.left == nil {
		if parent == nil {
			bst.root = temp
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
	fmt.Println("Element deleted")
}

func (bst *bstTree) findMinElement() {
	temp := bst.root
	minElement := temp.left
	for minElement.left != nil {
		minElement = minElement.left
	}
	fmt.Println("Minimum element in this tree is : ", minElement.data)
}

func (bst *bstTree) findMaxElement() {
	temp := bst.root
	maxElement := temp.right
	for maxElement.right != nil {
		maxElement = maxElement.right
	}
	fmt.Println("Max element in this tree is : ", maxElement.data)
}

func (bst *bstTree) printPreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d\n", node.data)
	bst.printPreOrder(node.left)
	bst.printPreOrder(node.right)

}
func (bst *bstTree) printPostOrder(node *Node) {
	if node == nil {
		return
	}
	bst.printPostOrder(node.left)
	bst.printPostOrder(node.right)
	fmt.Printf("%d\n", node.data)

}
func (bst *bstTree) printInOrder(node *Node) {
	if node == nil {
		return
	}

	bst.printInOrder(node.left)
	fmt.Printf("%d\n", node.data)
	bst.printInOrder(node.right)

}

func main() {
	BST := bstTree{}
	for true {
		var choice int
		fmt.Println("\nEnter any option : ")
		fmt.Println("1.Insert")
		fmt.Println("2.Inorder")
		fmt.Println("3.Preorder")
		fmt.Println("4.Postorder")
		fmt.Println("5.delete")
		fmt.Println("6.search")
		fmt.Println("7.Min Element")
		fmt.Println("8.Max Element")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			// var num int
			// fmt.Println("Enter the number : ")
			// fmt.Scan(&num)
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

		case 2:
			BST.printInOrder(BST.root)
		case 3:
			BST.printPreOrder(BST.root)
		case 4:
			BST.printPostOrder(BST.root)
		case 5:
			BST.delete(45)
		case 6:
			BST.contains(600)
		case 7:
			BST.findMinElement()
		case 8:
			BST.findMaxElement()

		}

	}

}
