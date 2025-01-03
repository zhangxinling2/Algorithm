package main

func numTrees(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		for j := i; j > 0; j-- {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
