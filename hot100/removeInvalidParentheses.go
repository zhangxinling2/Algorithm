package main

func removeInvalidParentheses(s string) []string {
	lRemove, rRemove := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lRemove++
		} else if s[i] == ')' {
			if lRemove > 0 {
				lRemove--
			} else {
				rRemove++
			}
		}
	}
	ans := []string{}
	dfs(&ans, s, 0, 0, 0, lRemove, rRemove)
	return ans
}

func dfs(ans *[]string, s string, lCount, rCount, start, lRemove, rRemove int) {
	if lRemove == 0 && rRemove == 0 {
		if valid(s) {
			*ans = append(*ans, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if s[i] == '(' {
			lCount++
		}
		if s[i] == ')' {
			rCount++
		}
		if i > start && s[i] == s[i-1] {
			continue
		}
		if s[i] == '(' && lRemove > 0 {
			dfs(ans, s[:i]+s[i+1:], lCount-1, rCount, i, lRemove-1, rRemove)
		}
		if s[i] == ')' && rRemove > 0 {
			dfs(ans, s[:i]+s[i+1:], lCount, rCount-1, i, lRemove, rRemove-1)
		}
		if rCount > lCount {
			break
		}
	}
}
func valid(s string) bool {
	cnt := 0
	for _, v := range s {
		if v == '(' {
			cnt++
		} else if v == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
func removeInvalidParentheses2(s string) []string {
	res := []string{}
	queue := []string{s}
	visited := map[string]bool{}
	visited[s] = true
	fonud := false
	for len(queue) > 0 {
		var level []string
		for _, str := range queue {
			if valid(str) {
				res = append(res, str)
				fonud = true
			}
			if !fonud {
				for i := 0; i < len(str); i++ {
					if str[i] != '(' && str[i] != ')' {
						continue
					}
					next := str[:i] + str[i+1:]
					if _, ok := visited[next]; !ok {
						level = append(level, next)
						visited[next] = true
					}
				}
			}
		}
		if fonud {
			break
		}
		queue = level
	}
	return res
}
