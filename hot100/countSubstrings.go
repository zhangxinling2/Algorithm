package main

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	count := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(dp); i++ {
		dp[i][i] = true
		count++
	}
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] == s[i] {
				if i-1 < j+1 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}
			if dp[j][i] {
				count++
			}
		}

	}
	return count
}

// 中心扩散法
func countSubstrings2(s string) int {
	n := len(s)
	// 初始化结果为0
	res := 0

	// 左右指针从中心扩展并比较
	var extend func(int, int)
	// left:左指针
	// right:右指针
	extend = func(left int, right int) {
		// 在左右指针不越界的情况下进行比较，直到左右指针字符不相等跳出循环为止
		for left >= 0 && right <= n-1 {
			if s[left] == s[right] {
				// 如果相等，则说明[left,right]范围内是回文子串，将结果+1
				res++
				// 继续以[left,right]为中心向外扩展判断[left-1,right+1]是否是回文子串
				left--
				right++
			} else {
				// 如果不相等，则说明[left,right]范围内不是回文子串
				break
			}
		}
	}

	// 遍历整个字符串，分别尝试以每个位置为中心进行扩展，来找到字符串中的所有回文子串
	for i := 0; i <= n-1; i++ {
		// 当回文字符串的长度为奇数时，则该回文字符串只能由 1个字符组成的中心 扩展而来，即以该字符为中心
		extend(i, i)
		// 当回文字符串的长度为偶数时，则该回文字符串只能由 2个字符组成的中心 扩展而来，即以该字符与其下一个字符一起组成中心
		extend(i, i+1)
	}

	// 返回最长回文子串结果
	return res
}
