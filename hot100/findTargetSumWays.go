package main

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if math.Abs(float64(target)) > float64(sum) {
		return 0 // 此时没有方案
	}
	if (target+sum)%2 == 1 {
		return 0 // 此时没有方案
	}
	bagSize := (target + sum) / 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, bagSize+1)
	}
	if nums[0] <= bagSize {
		dp[0][nums[0]] = 1
	}
	dp[0][0] = 1
	var numZero float64
	for i := range nums {
		if nums[i] == 0 {
			numZero++
		}
		dp[i][0] = int(math.Pow(2, numZero))
	}

	// 以下遍历顺序行列可以颠倒
	for i := 1; i < len(nums); i++ { // 行，遍历物品
		for j := 0; j <= bagSize; j++ { // 列，遍历背包
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][bagSize]
}
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bag+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}
	return dp[bag]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
