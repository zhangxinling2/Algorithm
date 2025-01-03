package main

import "container/list"

func rightSideView(root *TreeNode) []int {
	arr := []int{}
	level := list.New()
	if root == nil {
		return nil
	}
	level.PushBack(root)
	var node *TreeNode
	for level.Len() > 0 {
		l := level.Len()
		for i := 0; i < l; i++ {
			node = level.Remove(level.Front()).(*TreeNode)
			if node.Left != nil {
				level.PushBack(node.Left)
			}
			if node.Right != nil {
				level.PushBack(node.Right)
			}
		}
		arr = append(arr, node.Val)
	}
	return arr
}
