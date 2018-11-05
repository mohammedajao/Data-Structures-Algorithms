package main

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) peek() int {
	return q.head.data
}

func (q *Queue) add(data int) {
	node := &Node{data: data}
	if q.tail != nil {
		q.tail.next = node // Change previous tail
	}
	q.tail = node // Set new tail
	if q.head == nil {
		q.head = node
	}
}

func (q *Queue) remove() int {
	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return data
}

func main() {

}
