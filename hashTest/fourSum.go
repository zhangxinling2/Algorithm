package main

import "test/SortFunc"

func fourSum(nums []int, target int) [][]int {
	nums = SortFunc.QuickSort(nums, 0, len(nums)-1)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if n1 > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if n2 > target-n1 {
				break
			}
			if j > 1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if n1+n2+n3+n4 > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}
