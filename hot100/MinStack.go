package main

// type MinStack struct {
// 	Stack  []int
// 	Min    []int
// 	Length int
// }

// func MinStackConstructor() MinStack {
// 	return MinStack{
// 		Min:    []int{},
// 		Stack:  []int{},
// 		Length: 0,
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	if this.Length == 0 {
// 		this.Stack = append(this.Stack, val)
// 		this.Min = append(this.Min, val)
// 		this.Length++
// 		return
// 	}
// 	if val < this.Min[this.Length-1] {
// 		this.Min = append(this.Min, val)
// 	} else {
// 		this.Min = append(this.Min, this.Min[this.Length-1])
// 	}
// 	this.Stack = append(this.Stack, val)
// 	this.Length++
// }

// func (this *MinStack) Pop() {
// 	this.Stack = this.Stack[:this.Length-1]
// 	this.Min = this.Min[:this.Length-1]
// 	this.Length--

// }

// func (this *MinStack) Top() int {
// 	return this.Stack[this.Length-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.Min[this.Length-1]
//}
type MinStack struct {
	Stack []int
	Min   int
}

func MinStackConstructor() MinStack {
	return MinStack{
		Stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 {
		this.Min = val
	}
	diff := val - this.Min
	this.Stack = append(this.Stack, diff)
	if diff < 0 {
		this.Min = val
	}
}

func (this *MinStack) Pop() {
	diff := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	if diff < 0 {
		this.Min = this.Min - diff
	}
}

func (this *MinStack) Top() int {
	diff := this.Stack[len(this.Stack)-1]
	if diff < 0 { // 栈顶为最小值
		return this.Min
	} else { // 栈顶非最小值
		return this.Min + diff // 依 diff = val - mn 推导
	}
}

func (this *MinStack) GetMin() int {
	return this.Min
}
