package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func minDepth(root *TreeNode) int {
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
