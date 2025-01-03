package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		g[b] = append(g[b], a)
		indeg[a]++
	}
	q := []int{}
	for i, x := range indeg {
		if x == 0 {
			q = append(q, i)
		}
	}
	cnt := 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		cnt++
		for _, v := range g[i] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return cnt == numCourses
}
