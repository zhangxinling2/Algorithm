package main

func candy(ratings []int) int {
	res := make([]int, len(ratings))
	res[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		} else {
			res[i] = 1
		}
	}
	for i := len(res) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = max(res[i], res[i+1]+1)
		}
	}
	candys := 0
	for i := 0; i < len(res); i++ {
		candys += res[i]
	}
	return candys
}
