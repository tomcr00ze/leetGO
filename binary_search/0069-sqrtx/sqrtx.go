package sqrtx

/*
Problem: 69. Sqrt(x)
Link: https://leetcode.com/problems/sqrtx/
Category: binary_search
Difficulty: Easy

Time Complexity: O(log x)
Space Complexity: O(1)
*/

// MySqrt computes and returns the square root of x rounded down to the nearest integer.
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	l, r := 1, x/2
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
