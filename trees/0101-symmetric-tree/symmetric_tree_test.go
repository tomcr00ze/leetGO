package symmetrictree

import "testing"

func TestIsSymmetric(t *testing.T) {
	// Symmetric tree: 1 -> Left: 2 (3, 4), Right: 2 (4, 3)
	symTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	// Asymmetric tree
	asymTree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}

	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Symmetric tree",
			root:     symTree,
			expected: true,
		},
		{
			name:     "Asymmetric tree",
			root:     asymTree,
			expected: false,
		},
		{
			name:     "Nil tree",
			root:     nil,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSymmetric(tt.root)
			if got != tt.expected {
				t.Errorf("IsSymmetric() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}
	for i := 0; i < b.N; i++ {
		IsSymmetric(root)
	}
}
