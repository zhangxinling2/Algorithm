package main

import "math"

func minFunc(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minAbs(root *TreeNode) int {
	min := math.MaxInt32
	var pre *TreeNode
	var minAbsFunc func(root *TreeNode)
	minAbsFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		minAbsFunc(root.Left)
		if pre != nil {
			min = minFunc(min, root.Val-pre.Val)
		}
		pre = root
		minAbsFunc(root.Right)
	}
	minAbsFunc(root)
	return min
}
