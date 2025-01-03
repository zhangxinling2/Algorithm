package main

func isSymmetric(root *TreeNode) bool {
	var symmetric func(left *TreeNode, right *TreeNode) bool
	symmetric = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return symmetric(left.Left, right.Right) && symmetric(right.Left, left.Right)
	}
	return symmetric(root.Left, root.Right)
}
