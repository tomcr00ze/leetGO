package plusone

/*
Problem: 66. Plus One
Link: https://leetcode.com/problems/plus-one/
Category: arrays_hashing
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1) auxiliary space (O(N) if all digits are 9)
*/

// PlusOne increments the large integer represented by the digits array by 1.
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
