package main

func maxProfit(prices []int) int {
	maxVal, i := 0, 0
	for k := 0; k < len(prices); k++ {
		if prices[k] < prices[i] {
			i = k
		} else {
			maxVal = max(prices[k]-prices[i], maxVal)
		}
	}
	return maxVal
}
func maxProfit2(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

// 包含冷冻期
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = dp[i-1][1]
	}
	return max(dp[len(prices)-1][2], dp[len(prices)-1][1])
}
