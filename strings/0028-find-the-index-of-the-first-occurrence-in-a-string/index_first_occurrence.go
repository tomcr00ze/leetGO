package indexfirstoccurrence

import "strings"

/*
Problem: 28. Find the Index of the First Occurrence in a String
Link: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
Category: strings
Difficulty: Easy

Time Complexity: O(N * M)
Space Complexity: O(1)
*/

// StrStr returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
