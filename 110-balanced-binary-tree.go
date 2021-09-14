package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return getLevel(root) >= 0
}

func getLevel(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftH := getLevel(node.Left)
	rightH := getLevel(node.Right)
	if leftH >= 0 && rightH >= 0 && abs1(leftH, rightH) <= 1 {
		return max(leftH, rightH) + 1
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs1(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
