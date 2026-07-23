package addbinary

import "strconv"

/*
Problem: 67. Add Binary
Link: https://leetcode.com/problems/add-binary/
Category: bit_manipulation
Difficulty: Easy

Time Complexity: O(max(N, M))
Space Complexity: O(max(N, M))
*/

// AddBinary returns the sum of two binary strings as a binary string.
func AddBinary(a string, b string) string {
	i, j, carry, res := len(a)-1, len(b)-1, 0, ""
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	return res
}
