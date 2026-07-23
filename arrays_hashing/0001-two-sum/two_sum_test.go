package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "No Solution",
			nums:     []int{1, 2, 3},
			target:   10,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 23, 42, 55, 67, 89, 100}
	target := 189
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
