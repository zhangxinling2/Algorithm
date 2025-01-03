package main

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1
	index := len(nums) - 1
	for i <= j {
		m := nums[i] * nums[i]
		n := nums[j] * nums[j]
		if m >= n {
			res[index] = m
			i++
		} else {
			res[index] = n
			j--
		}
		index--
	}
	return res
}
