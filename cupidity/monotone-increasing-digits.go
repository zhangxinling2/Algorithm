package main

import (
	"math"
	"strconv"
)

func monotoneincreasingdigits(n int) int {
	nums := make([]int, 0)
	for n != 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i]
		j := i
		for j >= 1 && res[j] < res[j-1] {
			if res[j-1] == 9 && nums[j-1] != 9 {
				res[j] = 9
				break
			}
			res[j-1] = res[j-1] - 1
			res[j] = 9
			j--
		}
	}
	resN := 0
	for i := len(res) - 1; i >= 0; i-- {
		resN += res[i] * int(math.Pow(10, float64(len(res)-1-i)))
	}
	return resN
}
func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			for j := i; j < n; j++ { //后面的全部置为9
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
