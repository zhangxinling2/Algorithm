package main

func buildBinaryTreeFromNums(nums []int) *TreeNode {
	var buildFunc func(nums []int, i, j int) *TreeNode
	buildFunc = func(nums []int, i, j int) *TreeNode {
		if j < i {
			return nil
		}
		index := (j + i) / 2
		node := &TreeNode{
			Val:   nums[index],
			Left:  buildFunc(nums, i, index-1),
			Right: buildFunc(nums, index+1, j),
		}
		return node
	}
	return buildFunc(nums, 0, len(nums)-1)
}
