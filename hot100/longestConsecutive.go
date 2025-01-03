package main

func longestConsecutive(nums []int) int {
	index := map[int]bool{}
	ans := 0
	for i := 0; i < len(nums); i++ {
		index[nums[i]] = true
	}
	for i := range index {
		if index[i-1] {
			continue
		}
		j := i + 1
		for index[j] {
			j++
		}
		ans = max(ans, j-i)
	}
	return ans
}
