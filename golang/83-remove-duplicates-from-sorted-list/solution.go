package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head

	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
			continue
		}

		current = current.Next
	}

	return head
}
