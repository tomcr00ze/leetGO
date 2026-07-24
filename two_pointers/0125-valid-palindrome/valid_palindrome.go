package validpalindrome

import "strings"

/*
Problem: 125. Valid Palindrome
Link: https://leetcode.com/problems/valid-palindrome/
Category: two_pointers
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// IsPalindrome returns true if s is a palindrome after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	isAlnum := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
	}

	for l < r {
		for l < r && !isAlnum(s[l]) {
			l++
		}
		for l < r && !isAlnum(s[r]) {
			r--
		}
		if l < r && s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
