package main

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)]
}
func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
func rob2(root *TreeNode) int {
	var backtracking func(node *TreeNode) []int
	backtracking = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		left := backtracking(node.Left)
		right := backtracking(node.Right)
		notRob := max(left[0], left[1]) + max(right[0], right[1])
		rob := node.Val + left[0] + right[0]
		return []int{notRob, rob}
	}
	res := backtracking(root)
	return max(res[0], res[1])
}
