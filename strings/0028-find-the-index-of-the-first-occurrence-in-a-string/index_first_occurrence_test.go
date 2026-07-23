package indexfirstoccurrence

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Example 1",
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			name:     "Example 2",
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			name:     "Found in middle",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrStr(tt.haystack, tt.needle)
			if got != tt.expected {
				t.Errorf("StrStr(%q, %q) = %d; want %d", tt.haystack, tt.needle, got, tt.expected)
			}
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	haystack := "supercalifragilisticexpialidocious"
	needle := "expiali"
	for i := 0; i < b.N; i++ {
		StrStr(haystack, needle)
	}
}
