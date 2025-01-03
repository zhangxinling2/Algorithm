package main

func convertBST(root *TreeNode) *TreeNode {
	var pre *TreeNode
	if root == nil {
		return nil
	}
	var convertFunc func(root *TreeNode)
	convertFunc = func(root *TreeNode) {
		if root == nil {
			return 
		}
		convertFunc(root.Right)
		if pre == nil {
			pre = root
		} else {
			root.Val = root.Val + pre.Val
			pre = root
		}
		convertFunc(root.Left)
	}
	convertFunc(root)
	return root
}
