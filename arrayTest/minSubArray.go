package main

import "math"

func minSubArray(nums []int, target int) int {
	res, sum, i := math.MaxInt32, 0, 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if res > j-i+1 {
				res = j - i + 1
			}
			sum = sum - nums[i]
			i++
		}
	}
	return res
}
