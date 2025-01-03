package main

import "container/list"

func leftVal(root *TreeNode) int {
	maxDepth := 0
	res := 0
	var leftValFunc func(root *TreeNode, levelNum int)
	leftValFunc = func(root *TreeNode, levelNum int) {
		if root.Left == nil && root.Right == nil {
			if levelNum > maxDepth {
				res = root.Val
				maxDepth = levelNum
			}
			return
		}
		if root.Left != nil {
			leftValFunc(root.Left, levelNum+1)
		}
		if root.Right != nil {
			leftValFunc(root.Right, levelNum+1)
		}
		return
	}
	return res
}
func leftValII(root *TreeNode) int {
	tl := list.New()
	if root != nil {
		tl.PushBack(root)
	}
	res := root.Val
	var node *TreeNode
	for tl.Len() > 0 {
		length := tl.Len()
		for i := 0; i < length; i++ {
			node = tl.Remove(tl.Front()).(*TreeNode)
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				tl.PushBack(node.Left)
			}
			if node.Right != nil {
				tl.PushBack(node.Right)
			}
		}
	}
	return res
}
