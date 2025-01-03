package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	index := 0
	max := 0
	for i, num := range nums {
		if num > max {
			index = i
			max = num
		}
	}
	return buildBinaryTree(nums, index, 0, len(nums)-1)
}

func buildBinaryTree(nums []int, index, begin, end int) *TreeNode {
	if begin > end {
		return nil
	}
	if begin == end {
		return &TreeNode{Val: nums[begin]}
	}
	root := &TreeNode{Val: nums[index]}
	index1, index2 := begin, index
	max1, max2 := 0, 0
	for i := begin; i < index; i++ {
		if nums[i] > max1 {
			index1 = i
			max1 = nums[i]
		}
	}
	for i := index + 1; i <= end; i++ {
		if nums[i] > max2 {
			index2 = i
			max2 = nums[i]
		}
	}
	root.Left = buildBinaryTree(nums, index1, begin, index-1)
	root.Right = buildBinaryTree(nums, index2, index+1, end)
	return root
}
