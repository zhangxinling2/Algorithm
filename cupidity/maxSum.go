package main

import "math"

func maxSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	res := math.MinInt32
	for i := 0; i < len(dp); i++ {
		res = max(dp[i], res)
	}
	return res
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
