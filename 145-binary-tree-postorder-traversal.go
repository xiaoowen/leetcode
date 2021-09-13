package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)

	if root == nil {
		return vals
	}

	// left -> right -> node
	vals = append(vals, postorderTraversal(root.Left)...)
	vals = append(vals, postorderTraversal(root.Right)...)
	vals = append(vals, root.Val)
	return vals
}
