package singlenumber

/*
Problem: 136. Single Number
Link: https://leetcode.com/problems/single-number/
Category: bit_manipulation
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// SingleNumber finds the single element that appears only once using bitwise XOR.
func SingleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
