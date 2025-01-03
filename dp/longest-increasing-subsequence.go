package main

func longestIncreasingSubsequence(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}
func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	res := 0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
func findLengthII(nums1 []int, nums2 []int) int {
	n, m, res := len(nums1), len(nums2), 0
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := m; j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0 // 注意这里不相等要赋值为0，供下一层使用
			}
			res = max(res, dp[j])
		}
	}
	return res
}
