package main

func largestRectangleArea(heights []int) int {
	// 数组头部加入0
	heights = append([]int{0}, heights...)
	// 数组尾部加入0
	heights = append(heights, 0)
	stack := make([]int, 0)
	sum := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//左侧有低于index高度的柱子
			if len(stack) > 0 {
				sum = max(sum, (i-stack[len(stack)-1]-1)*heights[index])
			} else {
				//左侧没有低于index高度
				sum = max(sum, heights[i])
			}
		}
		stack = append(stack, i)
	}
	return sum
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
