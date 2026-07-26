package mergesortedarray

/*
Problem: 88. Merge Sorted Array
Link: https://leetcode.com/problems/merge-sorted-array/
Category: two_pointers
Difficulty: Easy

Time Complexity: O(M + N)
Space Complexity: O(1)
*/

// Merge merges nums2 into nums1 as one sorted array in-place.
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
