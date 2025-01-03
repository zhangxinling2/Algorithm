package main

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)]
}
func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
