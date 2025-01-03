package main

func wiggleMaxLength(nums []int) int {
	//if len(nums) < 2 {
	//	return len(nums)
	//}
	//path := []int{}
	//path = append(path, nums[0])
	//for i := 1; i < len(nums); {
	//	k := len(path) - 1
	//	if i < len(nums) && nums[i] == path[k] {
	//		i++
	//		continue
	//	}
	//	for i < len(nums) && k > 0 && path[k] > path[k-1] && path[k] <= nums[i] {
	//		i++
	//	}
	//	for i < len(nums) && k > 0 && path[k] < path[k-1] && path[k] >= nums[i] {
	//		i++
	//	}
	//	if i == len(nums) {
	//		break
	//	}
	//	path = append(path, nums[i])
	//}
	//return len(path)
	if len(nums) < 2 {
		return len(nums)
	}
	ans := 1
	prevDiff, curDiff := 0, 0
	for i := 1; i < len(nums); i++ {
		curDiff = nums[i] - nums[i-1]
		if curDiff > 0 && prevDiff <= 0 || curDiff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = curDiff
		}
	}
	return ans
}
