package twosum

/*
Problem: 1. Two Sum
Link: https://leetcode.com/problems/two-sum/
Category: arrays_hashing
Difficulty: Easy

Time Complexity: O(N) - single pass hash map traversal
Space Complexity: O(N) - storing elements in map
*/

// TwoSum finds indices of the two numbers such that they add up to target.
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}
