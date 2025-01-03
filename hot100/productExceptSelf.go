package main

//超出内存限制
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	var backtracking func(path []int, start int)
	paths := make([][]int, 0)
	backtracking = func(path []int, start int) {
		if len(path) == len(nums)-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}
		if len(path)+len(nums)-start < len(nums)-1 {
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(path, i+1)
			path = path[:len(path)-1]
		}
	}

	path := make([]int, 0)
	backtracking(path, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		val := 1
		for _, v := range paths[i] {
			val = val * v
		}
		res[len(nums)-1-i] = val
	}
	return res
}
func productExceptSelf2(nums []int) []int {
	L := make([]int, len(nums))
	R := make([]int, len(nums))
	L[0] = 1
	R[len(nums)-1] = 1
	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		L[i] = nums[i-1] * L[i-1]
		R[j] = nums[j+1] * R[j+1]
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = L[i] * R[i]
	}
	return res
}
func productExceptSelf3(nums []int) []int {
	pre, suf := 1, 1
	res := make([]int, len(nums))
	res[0], res[len(nums)-1] = 1, 1
	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		pre, suf = pre*nums[i-1], suf*nums[j+1]
		if i < j {
			res[i], res[j] = pre, suf
		} else if i == j {
			res[i] = pre * suf
		} else {
			res[i], res[j] = res[i]*pre, res[j]*suf
		}
	}

	return res
}
