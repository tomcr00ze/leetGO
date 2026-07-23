package searchinsert

/*
Problem: 35. Search Insert Position
Link: https://leetcode.com/problems/search-insert-position/
Category: binary_search
Difficulty: Easy

Time Complexity: O(log N)
Space Complexity: O(1)
*/

// SearchInsert returns the index if the target is found. If not, return the index where it would be if it were inserted in order.
func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
