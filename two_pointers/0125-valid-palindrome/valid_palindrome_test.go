package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid palindrome sentence",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "Empty string after cleanup",
			input:    " ",
			expected: true,
		},
		{
			name:     "Alphanumeric with numbers",
			input:    "0P",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	s := "A man, a plan, a canal: Panama! Was it a car or a cat I saw?"
	for i := 0; i < b.N; i++ {
		IsPalindrome(s)
	}
}
