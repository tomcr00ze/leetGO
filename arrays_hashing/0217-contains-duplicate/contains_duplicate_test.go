package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Contains duplicate",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "All unique elements",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Multiple duplicates",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Single element",
			nums:     []int{10},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsDuplicate(tt.nums)
			if got != tt.expected {
				t.Errorf("ContainsDuplicate(%v) = %v; want %v", tt.nums, got, tt.expected)
			}
		})
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1}
	for i := 0; i < b.N; i++ {
		ContainsDuplicate(nums)
	}
}
