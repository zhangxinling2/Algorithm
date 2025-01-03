package main

func isAnagram(s, t string) bool {
	record := make([]int, 26)
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		record[t[i]-'a']--
	}
	for _, i := range record {
		if i != 0 {
			return false
		}
	}
	return true
}
