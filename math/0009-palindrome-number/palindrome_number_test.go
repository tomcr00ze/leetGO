package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected bool
	}{
		{
			name:     "Positive palindrome",
			x:        121,
			expected: true,
		},
		{
			name:     "Negative number",
			x:        -121,
			expected: false,
		},
		{
			name:     "Trailing zero",
			x:        10,
			expected: false,
		},
		{
			name:     "Single digit zero",
			x:        0,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.x)
			if got != tt.expected {
				t.Errorf("IsPalindrome(%d) = %v; want %v", tt.x, got, tt.expected)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome(123454321)
	}
}
