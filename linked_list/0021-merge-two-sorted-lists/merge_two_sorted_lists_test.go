package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func buildLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func linkedListToSlice(head *ListNode) []int {
	res := []int{}
	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	return res
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "Example 1",
			l1:       []int{1, 2, 4},
			l2:       []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Both empty",
			l1:       []int{},
			l2:       []int{},
			expected: []int{},
		},
		{
			name:     "One empty list",
			l1:       []int{},
			l2:       []int{0},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := buildLinkedList(tt.l1)
			list2 := buildLinkedList(tt.l2)
			merged := MergeTwoLists(list1, list2)
			got := linkedListToSlice(merged)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeTwoLists() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l1 := buildLinkedList([]int{1, 3, 5, 7, 9})
		l2 := buildLinkedList([]int{2, 4, 6, 8, 10})
		MergeTwoLists(l1, l2)
	}
}
