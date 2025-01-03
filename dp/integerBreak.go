package main

func integerBreak(n int) int {
	maxR := 0
	//resS := make([]int, 0)
	var backtracking func(m int, startIndex int, res int)
	backtracking = func(m int, startIndex int, res int) {
		//fmt.Printf("m:%d,startIndex:%d,res:%d\m", m, startIndex, res)
		//fmt.Println(resS)
		if m == 0 {
			if res > maxR {
				maxR = res
			}
			return
		}
		if m < 0 {
			return
		}
		for i := startIndex; i <= m; i++ {
			if i == n {
				return
			}
			res *= i
			//resS = append(resS, i)
			backtracking(m-i, i, res)
			//resS = resS[:len(resS)-1]
			res /= i
		}
	}
	backtracking(n, 1, 1)
	return maxR
}
func integerBreak2(n int) int {
	/**
	  动态五部曲
	  1.确定dp下标及其含义
	  2.确定递推公式
	  3.确定dp初始化
	  4.确定遍历顺序
	  5.打印dp
	  **/
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
