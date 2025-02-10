package main

func maxCoins(nums []int) int {
	n := len(nums) + 2
	cache := make([][]int, n)
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, n)
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = -1
		}
	}
	arr := make([]int, n)
	for i := 0; i < len(nums); i++ {
		arr[i+1] = nums[i]
	}
	arr[0], arr[n-1] = 1, 1
	var solve func(l, r int) int
	solve = func(l, r int) int {
		if l+2 > r {
			return 0
		} else if l+2 == r {
			return arr[l] * arr[l+1] * arr[r]
		}
		if cache[l][r] != -1 {
			return cache[l][r]
		}
		res := 0
		k := l + 1
		for k <= r-1 {
			left := solve(l, k)
			right := solve(k, r)
			res = max(res, left+right+arr[r]*arr[k]*arr[l])
			k++
		}
		cache[l][r] = res
		return res
	}
	return solve(0, n-1)
}
func maxCoins2(nums []int) int {
	n := len(nums) + 2
	arr := make([]int, n)
	arr[0], arr[n-1] = 1, 1
	for i := 0; i < len(nums); i++ {
		arr[i+1] = nums[i]
	}
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+arr[k]*arr[i]*arr[j])
			}
		}
	}
	return dp[0][n-1]
}
