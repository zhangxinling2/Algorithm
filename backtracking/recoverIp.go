package main

import (
	"strconv"
	"strings"
)

//	func isValid(s string, start, end int) bool {
//		if start > end {
//			return false
//		}
//		if s[start] == '0' && start != end {
//			return false
//		}
//		num := 0
//		for i := start; i <= end; i++ {
//			if s[i] > '9' || s[i] < '0' {
//				return false
//			}
//			num = num*10 + s[i] - '0'
//			if num > 255 {
//				return false
//			}
//		}
//		return true
//	}
func recoverIp(s string) []string {
	res := []string{}
	path := []string{}
	var backtracking func(s string, startIndex int)
	backtracking = func(s string, startIndex int) {
		if len(path) == 4 {
			if startIndex == len(s) {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}
		for i := startIndex; i < len(s); i++ {
			if i != startIndex && s[startIndex] == '0' {
				break
			}
			str := s[startIndex : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				path = append(path, str)
				backtracking(s, i+1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	backtracking(s, 0)
	return res
}
