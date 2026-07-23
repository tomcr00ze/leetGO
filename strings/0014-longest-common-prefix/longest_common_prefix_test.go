package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "Common prefix flower",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "No common prefix",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "Single string",
			strs:     []string{"single"},
			expected: "single",
		},
		{
			name:     "Empty list",
			strs:     []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestCommonPrefix(tt.strs)
			if got != tt.expected {
				t.Errorf("LongestCommonPrefix(%v) = %q; want %q", tt.strs, got, tt.expected)
			}
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{"interspecies", "interstellar", "interstate"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix(strs)
	}
}
