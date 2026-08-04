package maxdepth

/*
Problem: 104. Maximum Depth of Binary Tree
Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/
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

// MaxDepth returns the maximum depth (number of nodes along longest path from root to leaf) of a binary tree.
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := MaxDepth(root.Left), MaxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
