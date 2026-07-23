package besttimestock

/*
Problem: 121. Best Time to Buy and Sell Stock
Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
Category: sliding_window
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// MaxProfit returns the maximum profit achievable from buying and selling stock once.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, maxProf := prices[0], 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProf {
			maxProf = p - minPrice
		}
	}
	return maxProf
}
