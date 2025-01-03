package main

func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
func isHappy(n int) bool {
	r := make(map[int]struct{})
	for n != 1 {
		n = getSum(n)
		if _, ok := r[n]; ok {
			return false
		} else {
			r[n] = struct{}{}
		}
	}
	return true
}
