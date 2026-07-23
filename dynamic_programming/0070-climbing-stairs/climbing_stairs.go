package climbingstairs

/*
Problem: 70. Climbing Stairs
Link: https://leetcode.com/problems/climbing-stairs/
Category: dynamic_programming
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// ClimbStairs returns the number of distinct ways to climb n stairs taking 1 or 2 steps.
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
