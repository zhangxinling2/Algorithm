package main

import (
	"math"
)

func perfect_squares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 1; j <= int(math.Sqrt(float64(n))); j++ {
			if i-j*j >= 0 && dp[i-j*j] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}
	return dp[n]
}
