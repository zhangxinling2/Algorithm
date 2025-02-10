package main

func lengthOfLIS(nums []int) int {
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

func lengthOfLIS2(nums []int) int {
	stack := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			replaceStack(stack, nums[i])
		}
	}
	return len(stack)
}
func replaceStack(stack []int, num int) {
	m, n := 0, len(stack)-1
	for m < n {
		mid := (m + n) / 2
		if stack[mid] >= num {
			n = mid
		} else {
			m = mid + 1
		}
	}
	stack[m] = num
}
