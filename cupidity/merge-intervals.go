package main

import "sort"

func mergeintervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			tmp := make([]int, 2)
			tmp[0] = start
			tmp[1] = end
			res = append(res, tmp)
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	tmp := make([]int, 2)
	tmp[0] = start
	tmp[1] = end
	res = append(res, tmp)
	return res
}
