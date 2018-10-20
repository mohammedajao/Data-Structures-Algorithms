package main /*
We will be implementing a Linked List.

- Data structure representing a collection of continguous elements
- Each node/element holds a reference to the next contingent element
- We have a head for a LinkedList.
- Head and tail for double LinkedList

- We will implement a LinkedList and Double LinkedList
*/

// Below we have a Doubly Linked List
// Golang's package for this is container/list

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}

	n = l.Last()
	for {
		println(n.value)
		n = n.Prev()
		if n == nil {
			break
		}
	}
}
