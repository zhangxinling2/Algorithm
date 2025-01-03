package main

func canCompleteCircuit(gas, cost []int) int {
	sum := 0
	c := 0
	index := 0
	cover := 0
	for i := 0; i < len(gas); i++ {
		c = gas[i] - cost[i]
		sum += c
		cover += c
		if cover < 0 {
			cover = 0
			index = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	return index
}
