package main

func generateMatrix(val int) [][]int {
	res := make([][]int, val)
	for i := 0; i < val; i++ {
		res[i] = make([]int, val)
	}
	startx, starty := 0, 0
	loop := val / 2
	length := val - 1
	count := 1
	for loop > 0 {
		i, j := startx, starty
		for j = starty; j < starty+length; j++ {
			res[i][j] = count
			count++
		}
		for i = startx; i < startx+length; i++ {
			res[i][j] = count
			count++
		}
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		length -= 2
		loop--
	}
	if val%2 == 1 {
		mid := val / 2
		res[mid][mid] = count
	}
	return res
}
