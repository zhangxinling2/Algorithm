package main

import "sort"

func combineToTarget(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(candidates)
	var backtracking func(candidates []int, target, index int)
	backtracking = func(candidates []int, target, index int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			path = append(path, candidates[i])
			backtracking(candidates, target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	backtracking(candidates, target, 0)
	return res
}
func combineToTargetII(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(candidates))
	sort.Ints(candidates)

	var backtracking func(candidates []int, target, index int)
	backtracking = func(candidates []int, target, index int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			if target-candidates[i] < 0 {
				break
			}
			path = append(path, candidates[i])
			used[i] = true
			backtracking(candidates, target-candidates[i], i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking(candidates, target, 0)
	return res
}
