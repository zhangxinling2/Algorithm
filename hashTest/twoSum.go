package main

func twoSum(nums []int, tar int) []int {
	r := make(map[int]int)
	for i, num := range nums {
		r[num] = i
	}
	res := make([]int, 2)
	for k, v := range r {
		dec := tar - k
		if val, ok := r[dec]; ok {
			res[0] = v
			res[1] = val
			return res
		}
	}
	return nil
}
