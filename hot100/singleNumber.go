package main

func singleNumber(nums []int) int {
	numsMap := make(map[int]interface{})
	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			delete(numsMap, nums[i])
		} else {
			numsMap[nums[i]] = struct{}{}
		}
	}
	res := 0
	for k, _ := range numsMap {
		res = k
	}
	return res
}
