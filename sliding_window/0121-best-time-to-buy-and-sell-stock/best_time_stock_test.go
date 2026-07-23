package besttimestock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Profitable array",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Monotonically decreasing prices",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Single price",
			prices:   []int{5},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.expected {
				t.Errorf("MaxProfit(%v) = %d; want %d", tt.prices, got, tt.expected)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4, 8, 2, 9, 12, 4, 15, 3, 19}
	for i := 0; i < b.N; i++ {
		MaxProfit(prices)
	}
}
