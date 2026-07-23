package lengthoflastword

import "strings"

/*
Problem: 58. Length of Last Word
Link: https://leetcode.com/problems/length-of-last-word/
Category: strings
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// LengthOfLastWord returns the length of the last word in the string.
func LengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	lastSpace := strings.LastIndex(s, " ")
	return len(s) - lastSpace - 1
}
