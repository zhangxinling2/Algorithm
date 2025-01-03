package main

import "container/list"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var reverseTree func(root *TreeNode)
	reverseTree = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		reverseTree(root.Left)
		reverseTree(root.Right)
	}
	reverseTree(root)
	return root
}
func invertTreeII(root *TreeNode) *TreeNode {
	if root != nil {
		return nil
	}
	ts := list.New()
	ts.PushBack(root)
	var node *TreeNode
	for ts.Len() > 0 {
		e := ts.Back().Value
		if e == nil {
			ts.Remove(ts.Back())
			node = ts.Remove(ts.Back()).(*TreeNode)
			continue
		}
		node = ts.Remove(ts.Back()).(*TreeNode)
		if node.Left != nil || node.Right != nil {
			node.Left, node.Right = node.Right, node.Left
		}
		if node.Right != nil {
			ts.PushBack(node.Right)
		}
		if node.Left != nil {
			ts.PushBack(node.Left)
		}
		ts.PushBack(node)
		ts.PushBack(nil)
	}
	return root
}
func invertTreeIII(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var reverseTree func(root *TreeNode)
	reverseTree = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		reverseTree(root.Left)
		reverseTree(root.Right)
	}
	reverseTree(root)
	return root
}
