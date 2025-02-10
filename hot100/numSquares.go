package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	minNum := math.MaxInt
	for i := 2; i < len(dp); i++ {
		for j := 1; j <= SqrtInt(i); j++ {
			minNum = min(minNum, dp[i-j*j]+1)
		}
		dp[i] = minNum
		minNum = math.MaxInt
	}
	return dp[n]
}
func SqrtInt(x int) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(ff)
}
