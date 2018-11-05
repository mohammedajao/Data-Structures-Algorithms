package main

// Queue: A collection of items that come out in FIFO(First-In-First-Out) order

type Node struct {
	value int
}

type Queue struct {
	data   []*Node
	size   int
	head   int
	tail   int
	length int
}

func NewQueue(size int) *Queue {
	return &Queue{
		data: make([]*Node, size),
		size: size,
	}
}

func (q *Queue) Enqueue(item int) {
	n := &Node{value: item}
	if (q.tail+1)%q.size == q.head {
		panic("Queue is full")
	}
	// if q.head == q.tail && q.length > 0 {
	// 	nodes := make([]*Node, len(q.data)+q.size)
	// 	copy(nodes, q.data[q.head:])
	// 	copy(nodes[len(q.data)-q.head:], q.data[:q.head])
	// 	q.head = 0
	// 	q.tail = len(q.data)
	// 	q.data = nodes
	// }
	q.data[q.tail] = n
	q.tail = (q.tail + 1) % q.size
	q.length++
}

func (q *Queue) Dequeue() *Node {
	if q.length == 0 {
		panic("Queue is empty")
	}
	node := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	q.length--
	return node
}

func (q *Queue) First() *Node {
	if q.length == 0 {
		panic("Queue is empty")
	}
	return q.data[q.head]
}

func (q *Queue) is_empty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

func main() {
	q := NewQueue(6)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	println(q.Dequeue().value)
	println(q.Dequeue().value)
	println(q.Dequeue().value)
	println(q.Dequeue().value)
	println(q.Dequeue().value)
	q.Enqueue(8)
	q.Enqueue(8)
	println(q.Dequeue().value)
	println(q.Dequeue().value)

}
