package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	vals := make([][]int, 0)
	if root == nil {
		return vals
	}

	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		levelVals := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)

		for _, node := range curLevel {
			levelVals = append(levelVals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		vals = append(vals, levelVals)
		// reverse
		if len(vals)%2 == 0 {
			reverseVals(levelVals)
		}
		curLevel = nextLevel
	}
	return vals
}

func reverseVals(in []int) {
	i := 0
	j := len(in) - 1
	for i < j {
		in[i], in[j] = in[j], in[i]
		j--
		i++
	}
}
