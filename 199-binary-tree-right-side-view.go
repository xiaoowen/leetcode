package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}

	curLevels := []*TreeNode{root}
	for len(curLevels) > 0 {
		vals = append(vals, curLevels[len(curLevels)-1].Val)

		nextLevels := make([]*TreeNode, 0)
		for _, node := range curLevels {
			if node.Left != nil {
				nextLevels = append(nextLevels, node.Left)
			}
			if node.Right != nil {
				nextLevels = append(nextLevels, node.Right)
			}
		}
		curLevels = nextLevels
	}

	return vals
}
