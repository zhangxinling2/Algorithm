package main

func binarySearch1(nums []int, target int) int {
	start, end, mid := 0, len(nums)-1, len(nums)/2
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
func binarySearch2(nums []int, target int) int {
	start, end, mid := 0, len(nums), len(nums)/2
	for start < end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}
