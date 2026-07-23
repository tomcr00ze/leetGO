package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     "Zero",
			x:        0,
			expected: 0,
		},
		{
			name:     "One",
			x:        1,
			expected: 1,
		},
		{
			name:     "Four",
			x:        4,
			expected: 2,
		},
		{
			name:     "Eight (round down)",
			x:        8,
			expected: 2,
		},
		{
			name:     "Large number",
			x:        2147395600,
			expected: 46340,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MySqrt(tt.x)
			if got != tt.expected {
				t.Errorf("MySqrt(%d) = %d; want %d", tt.x, got, tt.expected)
			}
		})
	}
}

func BenchmarkMySqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MySqrt(2147395600)
	}
}
