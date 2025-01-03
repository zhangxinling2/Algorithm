package main

func getZeroAndOne(str string) (int, int) {
	zeroN := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			zeroN++
		}
	}
	return zeroN, len(str) - zeroN
}
func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		zeroN, oneN := getZeroAndOne(strs[i])
		for j := m; j >= zeroN; j-- {
			for k := n; k >= oneN; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroN][k-oneN]+1)
			}
		}
	}
	return dp[m][n]
}
