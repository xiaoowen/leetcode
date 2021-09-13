package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)

	if root != nil {
		vals = append(vals, root.Val)
		vals = append(vals, preorderTraversal(root.Left)...)
		vals = append(vals, preorderTraversal(root.Right)...)
	}
	return vals
}
