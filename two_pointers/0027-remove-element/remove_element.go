package removeelement

/*
Problem: 27. Remove Element
Link: https://leetcode.com/problems/remove-element/
Category: two_pointers
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// RemoveElement removes all occurrences of val in nums in-place.
func RemoveElement(nums []int, val int) int {
	k := 0
	for _, n := range nums {
		if n != val {
			nums[k] = n
			k++
		}
	}
	return k
}
