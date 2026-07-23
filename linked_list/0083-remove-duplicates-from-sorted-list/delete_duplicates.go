package deleteduplicates

/*
Problem: 83. Remove Duplicates from Sorted List
Link: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
Category: linked_list
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteDuplicates removes all duplicates from a sorted linked list.
func DeleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
