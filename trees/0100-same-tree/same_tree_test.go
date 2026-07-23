package sametree

import "testing"

func TestIsSameTree(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	t4 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	tests := []struct {
		name     string
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			name:     "Identical trees",
			p:        t1,
			q:        t2,
			expected: true,
		},
		{
			name:     "Different structure",
			p:        t3,
			q:        t4,
			expected: false,
		},
		{
			name:     "Both nil",
			p:        nil,
			q:        nil,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameTree(tt.p, tt.q)
			if got != tt.expected {
				t.Errorf("IsSameTree() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	for i := 0; i < b.N; i++ {
		IsSameTree(p, q)
	}
}
