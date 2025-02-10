package main

func generateParenthesis(n int) []string {
	res := []string{}
	var backtracking func(l, r int, s string)
	backtracking = func(l, r int, s string) {
		if l == 0 && r == 0 {
			res = append(res, s)
			return
		}
		if l > r {
			return
		}
		if l > 0 {
			backtracking(l-1, r, s+"(")
		}
		if r > 0 {
			backtracking(l, r-1, s+")")
		}
	}
	backtracking(n, n, "")
	return res
}
