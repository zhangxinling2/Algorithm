package main

func repeatedSubstringPattern(s string) bool {
	next := make([]int, len(s))
	i, j := 0, -1
	next[0] = -1
	bytes := []byte(s)
	for i < len(bytes)-1 {
		if j == -1 || bytes[j] == bytes[i] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	if next[len(bytes)-1] != 0 && len(bytes)%(len(bytes)-next[len(bytes)-1]+1) == 0 {
		return true
	}
	return false
}
