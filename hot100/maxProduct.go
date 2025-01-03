package main

func maxProduct(nums []int) int {
	maxVal := nums[0]
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			dp[i][0] = min(nums[i]*dp[i-1][0], nums[i])
			dp[i][1] = max(nums[i]*dp[i-1][1], nums[i])

		} else {
			dp[i][0] = min(nums[i]*dp[i-1][1], nums[i])
			dp[i][1] = max(nums[i]*dp[i-1][0], nums[i])
		}
		if dp[i][1] > maxVal {
			maxVal = dp[i][1]
		}
	}
	return maxVal
}
