package searchinsert

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Target present",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Target absent in middle",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Target larger than all",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "Target smaller than all",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInsert(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("SearchInsert(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	for i := 0; i < b.N; i++ {
		SearchInsert(nums, 14)
	}
}
