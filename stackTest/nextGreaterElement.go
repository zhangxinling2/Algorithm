package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	index := make(map[int]int, len(nums2))
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			index[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		v, ok := index[nums1[i]]
		if !ok {
			res[i] = -1
			continue
		}
		res[i] = v
	}
	return res
}
func nextGreaterElementII(nums []int) []int {
	nums = append(nums, nums...)
	res := make([]int, len(nums))

	//存储下标
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(stack); i++ {
		res[stack[i]] = -1
	}
	return res[:len(nums)/2]
}
