package main

func canConstruct(ransomNote, magazine string) bool {
	res := make([]int, 26)
	for i := 0; i < len(ransomNote); i++ {
		res[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		res[magazine[i]-'a']--
	}
	for _, v := range res {
		if v > 0 {
			return false
		}
	}
	return true
}
