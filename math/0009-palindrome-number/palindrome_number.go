package palindromenumber

/*
Problem: 9. Palindrome Number
Link: https://leetcode.com/problems/palindrome-number/
Category: math
Difficulty: Easy

Time Complexity: O(log10 N)
Space Complexity: O(1)
*/

// IsPalindrome returns true if x is a palindrome integer.
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}
