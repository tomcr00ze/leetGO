package symmetrictree

/*
Problem: 101. Symmetric Tree
Link: https://leetcode.com/problems/symmetric-tree/
Category: trees
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(H) where H is tree height
*/

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSymmetric checks if a binary tree is a mirror of itself around its center.
func IsSymmetric(root *TreeNode) bool {
	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil || r == nil {
			return l == r
		}
		return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
	}
	return root == nil || check(root.Left, root.Right)
}
