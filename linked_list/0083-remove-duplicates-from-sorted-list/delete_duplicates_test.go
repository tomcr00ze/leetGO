package deleteduplicates

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

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.input)
			result := DeleteDuplicates(head)
			got := linkedListToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("DeleteDuplicates() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := buildLinkedList([]int{1, 1, 2, 3, 3, 4, 4, 4, 5})
		DeleteDuplicates(head)
	}
}
