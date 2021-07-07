package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

// PrintList : print all
func (n *ListNode) PrintList() {
	var node *ListNode
	node = n

	fmt.Println("self", node.Val)

	for node.Next != nil {
		node = node.Next
		fmt.Println("next", node.Val)
	}
}

// ReverseList : reverse
func (n *ListNode) ReverseList() *ListNode {
	// write code here
	if n == nil {
		return nil
	}

	pHead := n
	var newHead *ListNode
	for pHead != nil {
		node := pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = node
	}

	return newHead
}

func main() {
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	n0 := &ListNode{
		Val:  0,
		Next: n1,
	}

	fmt.Println("origin ---")
	n0.PrintList()

	n2.ReverseList()
	fmt.Println("reverse ---")
	n2.PrintList()
}
