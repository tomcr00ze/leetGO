package inordertraversal

/*
Problem: 94. Binary Tree Inorder Traversal
Link: https://leetcode.com/problems/binary-tree-inorder-traversal/
Category: trees
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(N) auxiliary space for recursion stack
*/

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal returns the inorder traversal of a binary tree's nodes' values.
func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
