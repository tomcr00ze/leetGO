package validparentheses

/*
Problem: 20. Valid Parentheses
Link: https://leetcode.com/problems/valid-parentheses/
Category: stack
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(N)
*/

// IsValid determines if the input string containing brackets is valid.
func IsValid(s string) bool {
	st := []rune{}
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		if match, ok := m[c]; ok {
			if len(st) == 0 || st[len(st)-1] != match {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return len(st) == 0
}
