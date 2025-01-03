package main

import "sort"

func biketochild(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	j := 0
	for i := 0; j < len(g) && i < len(s); i++ {
		if s[i] >= g[j] {
			res += 1
			j++
		} else {
			continue
		}
	}
	return res
}
