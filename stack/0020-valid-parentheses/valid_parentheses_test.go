package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple parentheses",
			input:    "()",
			expected: true,
		},
		{
			name:     "Multiple valid types",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "Mismatched type",
			input:    "(]",
			expected: false,
		},
		{
			name:     "Nested valid types",
			input:    "({[]})",
			expected: true,
		},
		{
			name:     "Unclosed opening bracket",
			input:    "([",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.input)
			if got != tt.expected {
				t.Errorf("IsValid(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	s := "({[()()]}){}[]()"
	for i := 0; i < b.N; i++ {
		IsValid(s)
	}
}
