package maxdepth

import "testing"

func TestMaxDepth(t *testing.T) {
	// Tree depth 3: 3 -> Left: 9, Right: 20 (15, 7)
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Depth 3 tree",
			root:     root,
			expected: 3,
		},
		{
			name:     "Nil tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepth(tt.root)
			if got != tt.expected {
				t.Errorf("MaxDepth() = %d; want %d", got, tt.expected)
			}
		})
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	for i := 0; i < b.N; i++ {
		MaxDepth(root)
	}
}
