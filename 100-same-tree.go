package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pvals := getVals(p)
	qvals := getVals(q)

	for i, v := range pvals {
		if v != qvals[i] {
			return false
		}
	}
	return true
}

func getVals(node *TreeNode) []int {
	vals := make([]int, 0)

	if node == nil {
		vals = append(vals, math.MaxInt32)
	} else {
		vals = append(vals, node.Val)
		vals = append(vals, getVals(node.Left)...)
		vals = append(vals, getVals(node.Right)...)
	}
	return vals
}
