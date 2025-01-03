package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		//if startIndex == len(nums) {
		//	tmp := make([]int, len(path))
		//	copy(tmp, path)
		//	res = append(res, tmp)
		//	return
		//}
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(nums, 0)
	res = append(res, []int{})
	return res
}
func subsets2(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := startIndex; i < len(nums); i++ {
			if i > 0 && used[i-1] == false && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true

			backtracking(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking(nums, 0)
	return res
}
