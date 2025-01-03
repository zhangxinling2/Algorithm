package main

func combine(n, k int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtracking func(n, k, startIndex int)
	backtracking = func(n, k, startIndex int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIndex; i <= n; i++ {
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(n, k, 1)
	return res
}
