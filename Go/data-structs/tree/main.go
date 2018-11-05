package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

type Tree struct {
	root   *Node
	length int
}

func NewTree() *Tree {
	return &Tree{}
}

func (n *Node) Insert(input int) {
	if input <= n.data {
		if n.left == nil {
			n.left = &Node{data: input}
		} else {
			n.left.Insert(input)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: input}
		} else {
			n.right.Insert(input)
		}
	}
}

func (n *Node) List() {
	if n.left != nil {
		n.left.List()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.List()
	}
}

func (t *Tree) Insert(input int) {
	if t.length == 0 {
		t.root = &Node{data: input}
	} else {
		t.root.Insert(input)
	}
	t.length += 1
}

func (t *Tree) ListNodes() {
	t.root.List()
}

func main() {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(8)
	tree.ListNodes()
}
