package main

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	res := []int{}
	for i, num := range nums {
		for len(q) > 0 && num > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-q[0] >= k {
			q = q[1:]
		}
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}
