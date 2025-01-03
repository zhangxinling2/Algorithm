package main

func tranversal(inorder []int, inorderBegin, inorderEnd int, postorder []int, postorderBegin, postorderEnd int) *TreeNode {
	if postorderBegin == postorderEnd {
		return nil
	}
	rootValue := postorder[postorderEnd-1]
	node := &TreeNode{Val: rootValue}
	if postorderEnd-postorderBegin == 1 {
		return node
	}
	var delimiterIndex int
	for delimiterIndex = inorderBegin; delimiterIndex < inorderEnd; delimiterIndex++ {
		if inorder[delimiterIndex] == rootValue {
			break
		}
	}
	leftInorderBegin := inorderBegin
	leftInorderEnd := delimiterIndex
	rightInorderBegin := delimiterIndex + 1
	rightInorderEnd := inorderEnd
	leftPostorderBegin := postorderBegin
	leftPostorderEnd := postorderBegin + delimiterIndex - inorderBegin
	rightPostorderBegin := postorderBegin + delimiterIndex - inorderBegin
	rightPostorderEnd := postorderEnd - 1
	node.Left = tranversal(inorder, leftInorderBegin, leftInorderEnd, postorder, leftPostorderBegin, leftPostorderEnd)
	node.Right = tranversal(inorder, rightInorderBegin, rightInorderEnd, postorder, rightPostorderBegin, rightPostorderEnd)
	return node
}
func buildTree(inorder, postorder []int) *TreeNode {
	return tranversal(inorder, 0, len(inorder), postorder, 0, len(postorder))
}
