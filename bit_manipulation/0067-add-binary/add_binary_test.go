package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "11 + 1",
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			name:     "1010 + 1011",
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			name:     "0 + 0",
			a:        "0",
			b:        "0",
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddBinary(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("AddBinary(%q, %q) = %q; want %q", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func BenchmarkAddBinary(b *testing.B) {
	a := "101010101010"
	bb := "110011001100"
	for i := 0; i < b.N; i++ {
		AddBinary(a, bb)
	}
}
