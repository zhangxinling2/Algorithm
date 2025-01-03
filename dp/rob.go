package main

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}
func robII(nums []int) int {
	if len(nums) <= 1 {
		return robWithoutCircle(nums)
	}
	res1 := robWithoutCircle(nums[:len(nums)-1])
	res2 := robWithoutCircle(nums[1:])
	if res1 > res2 {
		return res1
	}
	return res2
}
func robWithoutCircle(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)]
}
func robIII(cur *TreeNode) int {
	var backtracking func(cur *TreeNode) []int
	backtracking = func(cur *TreeNode) []int {
		if cur == nil {
			return []int{0, 0}
		}
		left := backtracking(cur.Left)
		right := backtracking(cur.Right)
		curRob := cur.Val + left[0] + right[0]
		notCurRob := max(left[0], left[1]) + max(right[0], right[1])
		return []int{notCurRob, curRob}
	}
	res := backtracking(cur)
	if res[0] > res[1] {
		return res[0]
	}
	return res[1]
}
