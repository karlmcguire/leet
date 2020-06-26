package main

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		a.Next = mergeTwoLists(a.Next, b)
		return a
	}
	b.Next = mergeTwoLists(a, b.Next)
	return b
}

func show(n *ListNode) string {
	b := &bytes.Buffer{}
	for i := n; i != nil; i = i.Next {
		fmt.Fprintf(b, "%d->", i.Val)
	}
	return b.String()[:b.Len()-2]
}

func main() {
	fmt.Printf("%v = 1->1->2->3->4->4\n", show(mergeTwoLists(
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
	)))
}
