package main

import "sort"

func overlappingIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			res++
		} else {
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}
	return len(intervals) - res
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
