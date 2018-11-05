package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	current := head.Next
	prev := head
	for current.Next != nil {
		future := current.Next
		current.Next = prev
		prev = current
		current = future
	}
	return head
}

func main() {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	println(reverseList(n1).Next.Next.Val)
}
