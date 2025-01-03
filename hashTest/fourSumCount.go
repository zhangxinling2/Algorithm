package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	abSum := make(map[int]int)
	for _, i := range nums1 {
		for _, i2 := range nums2 {
			abSum[i+i2]++
		}
	}
	count := 0
	for _, i := range nums3 {
		for _, i2 := range nums4 {
			if v, ok := abSum[-(i + i2)]; ok {
				if v > 0 {
					count++
					abSum[-(i+i2)]--
				}
			}
		}
	}
	return count
}
