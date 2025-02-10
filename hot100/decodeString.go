package main

import "strconv"

func decodeString(s string) string {
	stack := []rune{}
	for _, v := range s {
		if v == ']' {
			tmp := []rune{}
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == '[' {
					stack = stack[:len(stack)-1]
					break
				}
				tmp = append(tmp, stack[i])
				stack = stack[:len(stack)-1]
			}
			val := []rune{}
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == '[' || (stack[i] >= 'a' && stack[i] <= 'z') {
					break
				}

				val = append(val, stack[i])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
			}
			for i := 0; i < len(val)/2; i++ {
				val[i], val[len(val)-1-i] = val[len(val)-1-i], val[i]
			}
			num, _ := strconv.Atoi(string(val))
			for num > 0 {
				stack = append(stack, tmp...)
				num--
			}
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
