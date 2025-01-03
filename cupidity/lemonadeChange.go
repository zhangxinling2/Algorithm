package main

func lemonadeChange(bills []int) bool {
	capCharge := make([]int, 2)
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			capCharge[0] += 1
		}
		if bills[i] == 10 {
			capCharge[0] -= 1
			capCharge[1] += 1
		}
		if bills[i] == 20 {
			if capCharge[1] > 0 {
				capCharge[1] -= 1
				capCharge[0] -= 1
			} else {
				capCharge[0] -= 3
			}
		}
		if capCharge[0] < 0 {
			return false
		}
	}
	return true
}
