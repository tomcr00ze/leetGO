package removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		expectedK     int
		expectedArray []int
	}{
		{
			name:          "Example 1",
			nums:          []int{1, 1, 2},
			expectedK:     2,
			expectedArray: []int{1, 2},
		},
		{
			name:          "Example 2",
			nums:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedK:     5,
			expectedArray: []int{0, 1, 2, 3, 4},
		},
		{
			name:          "Empty slice",
			nums:          []int{},
			expectedK:     0,
			expectedArray: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			k := RemoveDuplicates(input)
			if k != tt.expectedK {
				t.Errorf("RemoveDuplicates() k = %d; want %d", k, tt.expectedK)
			}
			if k > 0 && !reflect.DeepEqual(input[:k], tt.expectedArray) {
				t.Errorf("RemoveDuplicates() slice = %v; want %v", input[:k], tt.expectedArray)
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9, 9}
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		RemoveDuplicates(input)
	}
}
