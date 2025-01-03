package main

func lastStone(stones []int) int {
	sums := 0
	for i := 0; i < len(stones); i++ {
		sums += stones[i]
	}
	sum := sums / 2
	dp := make([]int, sum+1)
	for i := 0; i < len(stones); i++ {
		for j := sum; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sums - dp[sum] - dp[sum]
}
