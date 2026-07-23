package sametree

/*
Problem: 100. Same Tree
Link: https://leetcode.com/problems/same-tree/
Category: trees
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(H) recursion stack where H is tree height
*/

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSameTree checks if two binary trees are structurally identical and have the same node values.
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
