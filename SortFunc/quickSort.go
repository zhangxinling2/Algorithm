package SortFunc

func QuickSort(nums []int, left, right int) []int {
	if left < right {
		par := partition(nums, left, right)
		QuickSort(nums, left, par-1)
		QuickSort(nums, par+1, right)
	}
	return nums
}
func partition(nums []int, left, right int) int {
	pivot := nums[left]
	i := left
	j := right
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
}
