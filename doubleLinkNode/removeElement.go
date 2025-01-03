package main

func removeElement(nums []int, val int) int {
	pre, last := 0, 0
	for last < len(nums) {
		for pre < len(nums) && nums[pre] != val {
			pre++
		}
		last = pre
		for last < len(nums) && nums[last] == nums[pre] {
			last++
		}
		if last == len(nums) {
			break
		}
		nums[pre], nums[last] = nums[last], nums[pre]
	}
	nums = nums[:pre]
	return pre
}
