package longestcommonprefix

import "strings"

/*
Problem: 14. Longest Common Prefix
Link: https://leetcode.com/problems/longest-common-prefix/
Category: strings
Difficulty: Easy

Time Complexity: O(N * M) where N is number of strings, M is length of first string
Space Complexity: O(1) auxiliary space
*/

// LongestCommonPrefix finds the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for _, s := range strs {
		for !strings.HasPrefix(s, pre) {
			pre = pre[:len(pre)-1]
			if len(pre) == 0 {
				return ""
			}
		}
	}
	return pre
}
