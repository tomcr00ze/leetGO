package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		val           int
		expectedK     int
		expectedArray []int
	}{
		{
			name:          "Example 1",
			nums:          []int{3, 2, 2, 3},
			val:           3,
			expectedK:     2,
			expectedArray: []int{2, 2},
		},
		{
			name:          "Example 2",
			nums:          []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:           2,
			expectedK:     5,
			expectedArray: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			k := RemoveElement(input, tt.val)
			if k != tt.expectedK {
				t.Errorf("RemoveElement() k = %d; want %d", k, tt.expectedK)
			}
			if !reflect.DeepEqual(input[:k], tt.expectedArray) {
				t.Errorf("RemoveElement() slice = %v; want %v", input[:k], tt.expectedArray)
			}
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 5, 2, 6, 2, 7}
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		RemoveElement(input, 2)
	}
}
