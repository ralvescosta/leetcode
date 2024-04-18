package main

import "testing"

func TestSolution(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	mergedList := MergeTwoLists(list1, list2)

	if mergedList != nil {
		t.Errorf("Merged list should be nil")
	}
}
