package main

func permutations(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	var backtracking func(nums []int)
	backtracking = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true

			backtracking(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking(nums)
	return res
}
func permutationsII(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	var backtracking func(nums []int)
	backtracking = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true

			backtracking(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking(nums)
	return res
}
