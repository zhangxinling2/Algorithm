package main

func trap(height []int) int {
	stack := make([]int, 0)
	weight := 0
	for i := 0; i < len(height); i++ {
		if len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				h := min(height[i], height[stack[len(stack)-1]]) - height[index]
				w := i - stack[len(stack)-1] - 1
				weight += h * w
			}
		}
		stack = append(stack, i)
	}
	return weight
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
