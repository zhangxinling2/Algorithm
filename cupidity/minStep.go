package main

func minStep(nums []int) int {
	cover := nums[0]
	tmpCover := 0
	step := make([]int, len(nums))
	i := 1
	k := 1
	step[0] = 0
	for i < len(nums) {
		for ; i <= cover; i++ {
			if i == len(nums) {
				break
			}
			step[i] = k
			tmpCover = max(tmpCover, i+nums[i])
		}
		k++
		cover = tmpCover
	}
	return step[len(nums)-1]
}
func minStep2(nums []int) int {
	cover := nums[0]
	step := 1
	tmpCover := 0
	for i := 0; i < len(nums)-1; i++ {
		tmpCover = max(i+nums[i], tmpCover)
		if i == cover {
			cover = tmpCover
			step++
		}
	}
	return step
}
