package inordertraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	// Tree: 1 -> Right: 2 -> Left: 3
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}

	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Example 1",
			root:     root,
			expected: []int{1, 3, 2},
		},
		{
			name:     "Nil root",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("InorderTraversal() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}},
	}
	for i := 0; i < b.N; i++ {
		InorderTraversal(root)
	}
}
