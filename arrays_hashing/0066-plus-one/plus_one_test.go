package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "Standard increment",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "Carry over single digit",
			digits:   []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "Nine carry over",
			digits:   []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "Multiple nines carry over",
			digits:   []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy slice so mutations don't interfere across runs
			input := make([]int, len(tt.digits))
			copy(input, tt.digits)
			got := PlusOne(input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("PlusOne(%v) = %v; want %v", tt.digits, got, tt.expected)
			}
		})
	}
}

func BenchmarkPlusOne(b *testing.B) {
	digits := []int{9, 9, 9, 9, 9, 9, 9, 9, 9}
	for i := 0; i < b.N; i++ {
		input := make([]int, len(digits))
		copy(input, digits)
		PlusOne(input)
	}
}
