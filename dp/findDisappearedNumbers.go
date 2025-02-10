package main

func findDisappearedNumbers(nums []int) []int {
	tmp := make([]int, len(nums))
	res := []int{}
	for i := 0; i < len(nums); i++ {
		tmp[nums[i]] = 1
	}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
func findDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); {
		idx := nums[i] - 1
		if nums[idx] == nums[i] {
			i++
			continue
		}
		nums[idx], nums[i] = nums[i], nums[idx]
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 { // 值与索引 不对应
			res = append(res, i+1)
		}
	}
	return res
}
