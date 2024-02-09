package collections

import (
	"fmt"
)

type OrderedList struct {
	assigner Assigner
	root     *Node
}

func NewOrderedList(assigner Assigner) *OrderedList {
	return &OrderedList{assigner: assigner}
}

type Node struct {
	getPriority func() float32
	value       interface{}
	left, right *Node
	height      int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func rotateRight(y *Node) *Node {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

func (ol *OrderedList) Insert(value interface{}) *Node {
	root := ol.root
	a := ol.assigner.AssignPriority(value)
	var insert func(node *Node, value interface{}, priority float32, assigner func() float32) *Node
	insert = func(node *Node, value interface{}, priority float32, assigner func() float32) *Node {
		if node == nil {
			return &Node{getPriority: assigner, value: value, height: 1}
		}

		if priority < node.getPriority() {
			node.left = insert(node.left, value, priority, assigner)
		} else if priority > node.getPriority() {
			node.right = insert(node.right, value, priority, assigner)
		} else {
			return node
		}

		node.height = 1 + max(height(node.left), height(node.right))

		balance := getBalance(node)

		if balance > 1 && priority < node.left.getPriority() {
			return rotateRight(node)
		}

		if balance < -1 && priority > node.right.getPriority() {
			return rotateLeft(node)
		}

		if balance > 1 && priority > node.left.getPriority() {
			node.left = rotateLeft(node.left)
			return rotateRight(node)
		}

		if balance < -1 && priority < node.right.getPriority() {
			node.right = rotateRight(node.right)
			return rotateLeft(node)
		}

		return node
	}

	return insert(root, value, a(), a)
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preOrder(node.left)
		preOrder(node.right)
	}
}
