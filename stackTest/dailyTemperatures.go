package main

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 初始化栈顶元素为第一个下标索引0
	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}
func dailyTemperatures2(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	idx := make([]int, len(temperatures))
	res := make([]int, len(temperatures))
	stack = append(stack, temperatures[0])
	idx = append(idx, 0)
	for i := 1; i < len(temperatures); i++ {
		if (len(stack)) == 0 {
			stack = append(stack, temperatures[i])
			idx = append(idx, i)
			continue
		}
		for temperatures[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			res[idx[len(idx)-1]] = i - idx[len(idx)-1]
			idx = idx[:len(idx)-1]
		}
		stack = append(stack, temperatures[i])
		idx = append(idx, i)
	}
	for i := 0; i < len(idx); i++ {
		res[idx[i]] = 0
	}
	return res
}
