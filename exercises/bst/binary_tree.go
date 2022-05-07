package bst

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (n *Node) Insert(data int) {
	if n.data > data {
		if n.left != nil {
			n.left.Insert(data)
			return
		}
		n.left = NewNode(data)
	} else if n.data < data {
		if n.right != nil {
			n.right.Insert(data)
			return
		}
		n.right = NewNode(data)
	}
}

func (n *Node) Contains(data int) *Node {
	if n.data == data {
		return n
	}
	if n.left == nil && n.right == nil {
		return nil
	}
	if n.data > data {
		return n.left.Contains(data)
	} else {
		return n.right.Contains(data)
	}
}

func Validate(n *Node, min, max *int) bool {
	if n == nil {
		return true
	}
	if (max == nil || (max != nil && n.data < *max)) && (min == nil || (min != nil && n.data > *min)) {
		if n.left != nil {
			return Validate(n.left, min, &n.data)
		}
		if n.right != nil {
			return Validate(n.right, &n.data, max)
		}
		return true
	}
	return false
}

func (n Node) Println() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.Println()
	}
	if n.right != nil {
		n.right.Println()
	}
}
