package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxVal := math.MinInt
	var TreeSum func(node *TreeNode) int
	TreeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftVal := TreeSum(node.Left)
		rightVal := TreeSum(node.Right)
		val := max(max(node.Val+leftVal, node.Val+rightVal), node.Val)
		maxVal = max(max(maxVal, val), node.Val+leftVal+rightVal)
		return val
	}
	TreeSum(root)
	return maxVal
}
