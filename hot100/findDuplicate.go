package main

func findDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		count := 0
		mid := (l + r) >> 1
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
