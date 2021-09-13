package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)

	if root != nil {
		vals = append(vals, inorderTraversal(root.Left)...)
		vals = append(vals, root.Val)
		vals = append(vals, inorderTraversal(root.Right)...)
	}
	return vals
}
