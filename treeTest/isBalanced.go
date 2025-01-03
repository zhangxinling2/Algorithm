package main

import "math"

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(getDepth(root.Left)-getDepth(root.Right))) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
