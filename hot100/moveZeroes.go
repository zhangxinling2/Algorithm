package main

func moveZeroes(nums []int) {
	left := []int{}
	right := []int{}
	m, n := 0, len(nums)-1
	for m <= n {
		for nums[m] != 0 && m <= n {
			left = append(left, nums[m])
			m++
		}
		for nums[n] == 0 && m <= n {
			n--
		}
		if nums[n] != 0 {
			right = append(right, nums[n])
		}
		nums[n] = nums[m]
		m++
		n--
	}
	m, n = 0, len(right)-1
	for m < n {
		right[m], right[n] = right[n], right[m]
		m++
		n--
	}
	copy(nums[:len(left)], left)
	copy(nums[len(left):len(left)+len(right)], right)
}
func moveZeroes2(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
