package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreeCameras(root *TreeNode) int {
	var res = 0
	var backtracking func(root *TreeNode) int
	backtracking = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		lv := backtracking(root.Left)
		rv := backtracking(root.Right)
		if lv == 2 && rv == 2 {
			return 0
		}
		if lv == 0 || rv == 0 {
			res++
			return 1
		}
		if lv == 1 || rv == 1 {
			return 2
		}

		return -1
	}
	if backtracking(root) == 0 {
		res++
	}
	return res
}
