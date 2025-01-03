package main

func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, bagweight+1)
	}
	//初始化
	for i := 0; i <= bagweight; i++ {
		if i > weight[0] {
			dp[0][i] = dp[0][i-weight[0]] + value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		for j := 1; j <= bagweight; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]
}
func test_1_wei_bag_problem1(weight, value []int, bagweight int) int {
	dp := make([]int, bagweight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagweight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagweight]
}
