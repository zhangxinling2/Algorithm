package main

func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
func canJump2(nums []int) bool {
	var backtracking func(nums []int, k int) bool
	backtracking = func(nums []int, k int) bool {
		if k > len(nums)-1 {
			return false
		}
		if k == len(nums)-1 {
			return true
		}
		for i := 0; i <= nums[k]; i++ {
			if i == 0 {
				continue
			}
			if backtracking(nums, k+i) {
				return true
			}
		}
		return false
	}
	return backtracking(nums, 0)

}
