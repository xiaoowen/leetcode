package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	leftVals := getVals(root.Left, "left")
	rightVals := getVals(root.Right, "right")
	if len(leftVals) != len(rightVals) {
		return false
	}

	for i, v := range leftVals {
		if v != rightVals[i] {
			return false
		}
	}
	return true
}

// 后序遍历
func getVals(node *TreeNode, orderType string) []int {
	vals := make([]int, 0)

	if node == nil {
		vals = append(vals, -1)
	} else {
		if orderType == "left" {
			vals = append(vals, getVals(node.Left, orderType)...)
			vals = append(vals, getVals(node.Right, orderType)...)
		} else {
			vals = append(vals, getVals(node.Right, orderType)...)
			vals = append(vals, getVals(node.Left, orderType)...)
		}
		vals = append(vals, node.Val)
	}
	return vals
}
