package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n-1]
}
func climbStairsII(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(cost)]
}

func climbStairsIII(n, m int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}
