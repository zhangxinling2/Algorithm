package main

func majorityElement(nums []int) int {
	curNum := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != curNum {
			cnt--
		} else {
			cnt++
		}
		if cnt == 0 {
			curNum = nums[i]
			cnt++
		}
	}
	return curNum
}
