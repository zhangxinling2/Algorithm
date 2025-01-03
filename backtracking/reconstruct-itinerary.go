package main

import "sort"

type pair struct {
	target string
	used   bool
}
type pairs []*pair

// Len implements sort.Interface.
func (p pairs) Len() int {
	return len(p)
}

// Less implements sort.Interface.
func (p pairs) Less(i int, j int) bool {
	return p[i].target < p[j].target
}

// Swap implements sort.Interface.
func (p pairs) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func reconstructitinerary(strs [][]string) []string {
	res := []string{}
	targets := make(map[string]pairs)
	for i := 0; i < len(strs); i++ {
		if targets[strs[i][0]] == nil {
			targets[strs[i][0]] = make(pairs, 0)
		}
		targets[strs[i][0]] = append(targets[strs[i][0]], &pair{
			target: strs[i][1],
			used:   false,
		})
	}
	for k, _ := range targets {
		sort.Sort(targets[k])
	}
	res = append(res, "JFK")
	var backtracking func() bool
	backtracking = func() bool {
		if len(res) == len(strs)+1 {
			return true
		}
		for _, pair := range targets[res[len(res)-1]] {
			if !pair.used {
				res = append(res, pair.target)
				pair.used = true
				if backtracking() {
					return true
				}
				res = res[:len(res)-1]
				pair.used = false
			}	
		}
		return false
	}
	backtracking()
	return res
}
