package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	case2(s[:3])
	PrintSliceStruct(&s)
}

// 更改元素
func case1(s []int) {
	s[0] = 1
	PrintSliceStruct(&s)
}
func case2(s []int) {
	s = append(s, 0)
	PrintSliceStruct(&s)
}

// 截取元素
func case3(s []int) {
	s = s[1 : len(s)-1]
	PrintSliceStruct(&s)
}

// 删除元素
func case4(s []int) {
	s1 := append(s[:1], s[2:]...)
	PrintSliceStruct(&s1)
}
func PrintSliceStruct(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("slice struct: %+v,slice is %v\n", ss, s)
}
