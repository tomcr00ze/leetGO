package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "III",
			input:    "III",
			expected: 3,
		},
		{
			name:     "LVIII",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "MCMXCIV",
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RomanToInt(tt.input)
			if got != tt.expected {
				t.Errorf("RomanToInt(%q) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	s := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		RomanToInt(s)
	}
}
