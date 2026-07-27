package removeduplicates

/*
Problem: 26. Remove Duplicates from Sorted Array
Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Category: two_pointers
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// RemoveDuplicates removes duplicates in-place such that each unique element appears only once.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
