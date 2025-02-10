package main

func countBits(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		res = append(res, calcBits(i))
	}
	return res
}
func calcBits(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}
