package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "Trailing spaces",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Example 3",
			input:    "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLastWord(tt.input)
			if got != tt.expected {
				t.Errorf("LengthOfLastWord(%q) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	s := "   fly me   to   the moon  "
	for i := 0; i < b.N; i++ {
		LengthOfLastWord(s)
	}
}
