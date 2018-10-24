package main /*
We will be implementing a Linked List.

- Data structure representing a collection of continguous elements
- Each node/element holds a reference to the next contingent element
- We have a head for a LinkedList.
- Head and tail for double LinkedList

- We will implement a LinkedList and Double LinkedList
*/

type List struct {
	head   *Node
	tail   *Node
	length int
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Find(value int) *Node {
	n := l.First()
	for {
		if n.value == value {
			return n
		}
		n = n.Next()
	}
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
	l.length += 1
}

func (l *List) Insert(value, position int) {
	node := &Node{value: value}
	n := l.First()
	count := 0
	if position == 0 {
		l.head = node
		node.next = n
		n.prev = node
		return
	} else if position == l.length-1 {
		node.prev = l.tail
		l.tail = node
		node.prev.next = node
		return
	}
	for {
		if count == position {
			node.next = n
			if n.prev != nil {
				node.prev = n.Prev()
				n.prev.next = node
			}
			n.prev = node
			break
		}
		n = n.Next()
		count += 1
		if n == nil {
			break
		}
	}
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
	// Test cases
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	l.Insert(7, 2)
	n := l.First()
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("-----")
	n = l.Last()
	// l.Insert(5, 0)
	for {
		println(n.value)
		n = n.Prev()
		if n == nil {
			break
		}
	}
}
