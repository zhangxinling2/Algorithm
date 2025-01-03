package main

func combineIII(n, k int) [][]int {
	res := [][]int{}
	path := []int{}
	sum := 0
	var backtracking func(n, k, startIndex int)
	backtracking = func(n int, k int, startIndex int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := startIndex; i <= 9; i++ {
			if 9-i+1 < k-len(path) || sum+i > n {
				break
			}
			sum += i
			path = append(path, i)
			backtracking(n, k, i+1)
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	backtracking(n, k, 1)
	return res
}
