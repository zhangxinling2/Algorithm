package main

func canPartition(nums []int) bool {
	sums := 0
	for i := 0; i < len(nums); i++ {
		sums += nums[i]
	}
	if sums%2 == 1 {
		return false
	}
	target := sums / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j-nums[i]] + nums[i]
		}
	}
	return dp[target] == target
}
