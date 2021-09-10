package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals1 := getLeafs(root1)
	vals2 := getLeafs(root2)

	if len(vals1) != len(vals2) {
		return false
	}

	for index, val := range vals1 {
		if val != vals2[index] {
			return false
		}
	}
	return true
}

// find leaf
func getLeafs(node *TreeNode) []int {
	vals := make([]int, 0)
	if node == nil {
		return vals
	}
	if node.Left == nil && node.Right == nil {
		vals = append(vals, node.Val)
		return vals
	}
	vals = append(vals, getLeafs(node.Left)...)
	vals = append(vals, getLeafs(node.Right)...)
	return vals
}
