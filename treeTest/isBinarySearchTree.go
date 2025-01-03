package main

func isBinarySearchTree(root *TreeNode) bool {
	res := []int{}
	var traversal func(root *TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		if root.Left != nil {
			traversal(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			traversal(root.Right)
		}
	}
	traversal(root)
	for i := 0; i < len(res)-1; i++ {
		j := i + 1
		if res[i] >= res[j] {
			return false
		}
	}
	return true
}
