package main

func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		used := make(map[int]bool, len(nums))
		for i := startIndex; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				backtracking(nums, i+1)
				path = path[:len(path)-1]
			}

		}
	}
	backtracking(nums, 0)
	return res
}
