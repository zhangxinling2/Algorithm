package main

func findNext(str string) []int {
	j := -1
	next := make([]int, len(str))
	next[0] = j
	for i := 1; i < len(str); i++ {
		for j >= 0 && str[i] != str[j+1] {
			j = next[j]
		}
		if str[i] == str[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}
func strStr(str1, str2 string) int {
	next := findNext(str2)
	j := -1
	for i := 0; i < len(str1); i++ {
		for j >= 0 && str2[j+1] != str1[i] {
			j = next[j]
		}
		if str1[i] == str2[j+1] {
			j++
		}
		if j == len(str2)-1 {
			return i - len(str2) + 1
		}
	}
	return -1
}
