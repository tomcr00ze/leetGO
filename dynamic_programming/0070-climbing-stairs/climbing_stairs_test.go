package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "2 stairs",
			n:        2,
			expected: 2,
		},
		{
			name:     "3 stairs",
			n:        3,
			expected: 3,
		},
		{
			name:     "5 stairs",
			n:        5,
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.n)
			if got != tt.expected {
				t.Errorf("ClimbStairs(%d) = %d; want %d", tt.n, got, tt.expected)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairs(45)
	}
}
