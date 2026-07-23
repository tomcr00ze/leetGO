package romantointeger

/*
Problem: 13. Roman to Integer
Link: https://leetcode.com/problems/roman-to-integer/
Category: strings
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// RomanToInt converts a Roman numeral string to an integer.
func RomanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			res -= m[s[i]]
		} else {
			res += m[s[i]]
		}
	}
	return res
}
