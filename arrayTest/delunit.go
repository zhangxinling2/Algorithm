package main

func delUnit(num []int, val int) []int {
	j := 0
	for i := 0; i < len(num); i++ {
		if num[i] != val {
			num[j] = num[i]
			j++
		}
	}
	return num[:j]
}
