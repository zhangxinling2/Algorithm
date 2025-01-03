package main

func insertIntoSearchTree(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	cur := root
	pre := root
	for cur != nil {
		pre = cur
		if p.Val < cur.Val {
			cur = cur.Left
		}
		if p.Val > cur.Val {
			cur = cur.Right
		}
	}
	if p.Val > pre.Val {
		pre.Right = p
	} else {
		pre.Left = p
	}
	return root
}
