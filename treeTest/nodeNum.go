package main

func nodeNum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftN, rightN := 0, 0
	l, r := root, root
	for l.Left != nil {
		leftN++
		l = l.Left
	}
	for r.Right != nil {
		rightN++
		r = r.Right
	}
	if leftN == rightN {
		return 2<<leftN - 1
	}
	return nodeNum(root.Left) + nodeNum(root.Right) + 1
}
