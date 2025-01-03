package main

func hasPath(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val-root.Val < 0 {
		return false
	} else if val-root.Val == 0 {
		if root.Left != nil || root.Right != nil {
			return false
		} else {
			return true
		}
	}
	return hasPath(root.Left, val-root.Val) || hasPath(root.Right, val-root.Val)
}
