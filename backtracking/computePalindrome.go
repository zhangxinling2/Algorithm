package main

func isPalindrome(s string) bool {
	l := len(s)
	i, j := 0, l-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func computePalindrome(s string) [][]string {
	res := [][]string{}
	path := []string{}
	var backtracking func(s string, startIndex int)
	backtracking = func(s string, startIndex int) {
		if startIndex >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				backtracking(s, i+1)
				path = path[:len(path)-1]
			}
		}
		return
	}
	backtracking(s, 0)
	return res
}
