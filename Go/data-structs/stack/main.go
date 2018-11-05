package main

import (
	"math"
)

type Node struct {
	value interface{}
}

type Stack struct {
	nodes  []*Node
	length int
}

func NewStack() *Stack {
	return &Stack{
		nodes:  make([]*Node, 0),
		length: 0,
	}
}

func (s *Stack) Push(item *Node) {
	s.nodes = append(s.nodes[:s.length], item)
	s.length++
}

func (s *Stack) Pop() *Node {
	if s.length == 0 {
		panic("Stack is empty")
	}
	s.length--
	return s.nodes[s.length]
}

func main() {
	s := NewStack()
	for i := 0; i < 10; i++ {
		n := &Node{value: math.Pow(float64(i), float64(2))}
		s.Push(n)
	}
	for {
		if s.length == 0 {
			break
		}
		println(s.Pop().value)
	}
}
