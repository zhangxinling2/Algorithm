package main

import (
	"fmt"
)

func coin_change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1000
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == 1000 {
		return -1
	}
	return dp[amount]
}

func coin_chang_ii(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
			fmt.Printf("coin:%d amount:%d func:%d\n", coins[i], j, dp[j])
			fmt.Println(dp)
		}
	}
	//for i := 0; i <= amount; i++ {
	//	for j := 0; j < len(coins); j++ {
	//		if i-coins[j] >= 0 {
	//			dp[i] += dp[i-coins[j]]
	//			fmt.Printf("coin:%d amount:%d func:%d\n", coins[j], i, dp[i])
	//			fmt.Println(dp)
	//		}
	//	}
	//}
	return dp[amount]
}
