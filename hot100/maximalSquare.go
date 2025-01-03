package main

import (
	"math"
)

func maximalSquare(matrix [][]byte) int {
	max := 0
	var isSquare func(m, n, col int) bool
	isSquare = func(m, n, col int) bool {
		for i := 0; i < col; i++ {
			if matrix[m+i][n+col] == 0 || matrix[m+col][n+i] == 0 {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				for k := 0; k < min(len(matrix[0])-j, len(matrix)-i); k++ {
					if !isSquare(i, j, k) {
						break
					}
					if ((k + 1) * (k + 1)) > max {
						max = (k + 1) * (k + 1)
					}
				}
			}
		}
	}
	return max
}
func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
func minThree(m, n, k int) int {
	min := 0
	if m < n {
		min = m
	} else {
		min = n
	}
	if k < min {
		min = k
	}
	return min
}
func mySqrt(x byte) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(ff)
}
func maximalSquare2(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = int(matrix[i][0])
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = int(matrix[0][i])
	}
	max := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			square := matrix[i-1][j-1]
			if square == 0 {
				dp[i][j] = 1
			}
			tmp := mySqrt(square)
			for k := 1; k <= tmp; k++ {
				if matrix[i][j-k] == 0 || matrix[i-k][j] == 0 {
					dp[i][j] = (k + 1) * (k + 1)
					if dp[i][j] > max {
						max = dp[i][j]
					}
					break
				}
				dp[i][j] = (k + 1) * (k + 1)
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}
func maximalSquare3(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = minThree(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max * max
}

// func main() {
// 	martix := [][]byte{
// 		[]byte{1, 0, 1, 0, 0},
// 		[]byte{1, 0, 1, 1, 1},
// 		[]byte{1, 1, 1, 1, 1},
// 		[]byte{1, 0, 0, 1, 0},
// 	}
// 	fmt.Println(maximalSquare3(martix))

// }
