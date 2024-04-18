package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{}
	cur := merged
	for list1 != nil || list2 != nil {
		if (list1 != nil && list2 != nil && list1.Val <= list2.Val) || list2 == nil {
			cur.Next = list1
			list1 = list1.Next
		} else if list2 != nil {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	return merged.Next
}
