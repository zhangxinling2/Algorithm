package main

//超时s ="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
func wordBreak(s string, wordDict []string) bool {
	var backtracking func(stringStart int) bool
	backtracking = func(stringStart int) bool {
		if stringStart == len(s) {
			return true
		}
		for i := 0; i < len(wordDict); i++ {
			l := len(wordDict[i])
			if stringStart+l > len(s) || s[stringStart:stringStart+l] != wordDict[i] {
				continue
			}
			if backtracking(stringStart + l) {
				return true
			}
		}
		return false
	}
	return backtracking(0)
}
func wordBreak2(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s))
	dp[0] = true
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
