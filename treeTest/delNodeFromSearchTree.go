package main

func delNodeFromSearchTree(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	if root.Val > key {
		root.Left = delNodeFromSearchTree(root.Left, key)
	} else {
		root.Right = delNodeFromSearchTree(root.Right, key)
	}
	return root
}
