package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber(tt.nums)
			if got != tt.expected {
				t.Errorf("SingleNumber(%v) = %d; want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2, 9, 9, 12, 12, 5, 5}
	for i := 0; i < b.N; i++ {
		SingleNumber(nums)
	}
}
