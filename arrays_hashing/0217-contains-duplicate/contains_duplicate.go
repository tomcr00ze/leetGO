package containsduplicate

/*
Problem: 217. Contains Duplicate
Link: https://leetcode.com/problems/contains-duplicate/
Category: arrays_hashing
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(N)
*/

// ContainsDuplicate returns true if any value appears at least twice in the array.
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}
