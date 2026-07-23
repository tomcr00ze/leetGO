package mergetwosortedlists

/*
Problem: 21. Merge Two Sorted Lists
Link: https://leetcode.com/problems/merge-two-sorted-lists/
Category: linked_list
Difficulty: Easy

Time Complexity: O(N + M)
Space Complexity: O(1)
*/

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two sorted linked lists into one sorted linked list.
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}
