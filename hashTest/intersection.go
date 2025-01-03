package main

func intersection(nums1 []int, nums2 []int) []int {
	sec := make(map[int]int)
	for _, i := range nums1 {
		sec[i] = 1
	}
	for _, i := range nums2 {
		if _, ok := sec[i]; ok {
			sec[i] = 0
		} else {
			continue
		}
	}
	res := make([]int, 0)
	for k, v := range sec {
		if v == 0 {
			res = append(res, k)
		}
	}
	return res
}
