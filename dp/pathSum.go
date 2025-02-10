package main

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
func pathSum2(root *TreeNode, targetSum int) int {
	m := map[int]int{}
	ans := 0
	m[0] = 1
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, prefixSum int) {
		if node == nil {
			return
		}
		prefixSum += node.Val
		ans += m[prefixSum-targetSum]
		m[prefixSum] += 1
		dfs(node.Left, prefixSum)
		dfs(node.Right, prefixSum)
		m[prefixSum] -= 1
	}
	dfs(root, 0)
	return ans
}
