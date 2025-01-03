package main

import "sort"

func splitSlice(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]int, sum+1)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[sum] == sum {
		return true
	}
	return false
}
