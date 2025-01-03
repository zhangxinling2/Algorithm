package main

func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	if dp[len(t)][len(s)] == len(s) {
		return true
	}
	return false
}
