package main

import "sort"

func burstBalloons(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	stepLen := points[0][1]
	res := 0
	for i := 0; i < len(points); i++ {
		if points[i][0] > stepLen {
			res++
			stepLen = points[i][1]
		}
		if points[i][1] < stepLen {
			stepLen = points[i][1]
		}
	}
	return res + 1
}
