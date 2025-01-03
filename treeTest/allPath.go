package main

func allPath(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	tmpRes := []int{}
	var getPath func(root *TreeNode)
	getPath = func(root *TreeNode) {
		tmpRes = append(tmpRes, root.Val)
		if root.Left == nil && root.Right == nil {
			tmp := []int{}
			tmp = append(tmp, tmpRes...)
			res = append(res, tmp)
			tmpRes = tmpRes[:len(tmpRes)-1]
			return
		}
		if root.Left != nil {
			getPath(root.Left)
		}
		if root.Right != nil {
			getPath(root.Right)
		}
		tmpRes = tmpRes[:len(tmpRes)-1]
	}
	getPath(root)
	return res
}
