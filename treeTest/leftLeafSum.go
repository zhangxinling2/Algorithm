package main

func leftLeafSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		if root.Right != nil {
			return leftLeafSum(root.Right) + root.Left.Val
		}
		return root.Left.Val
	}
	return leftLeafSum(root.Left) + leftLeafSum(root.Right)
}
