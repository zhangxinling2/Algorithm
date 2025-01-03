package main

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}
func merge(left []int, right []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for i < len(left) {
		res = append(res, left[i])
		i++
	}
	for j < len(right) {
		res = append(res, right[j])
		j++
	}
	return res
}
func quickSort(arr []int, left, right int) []int {
	if left < right {
		parIndex := partition(arr, left, right)
		quickSort(arr, left, parIndex-1)
		quickSort(arr, parIndex+1, right)
	}
	return arr
}
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		//如果当前值比较大，那么i向后移动到小值与index交换，最后index-1的位置就是中间位置
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
