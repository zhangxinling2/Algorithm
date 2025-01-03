package main

import "fmt"

func createLargestLand() {
	var (
		m int
		n int
	)
	fmt.Scanln(&m, &n)
	type point struct {
		x int
		y int
	}
	directions := []point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	visited := make([][]bool, m)
	acr := make(map[int]int, 0)
	index := 0
	graph := make([][]int, m)
	for i := 0; i < len(graph); i++ {
		visited[i] = make([]bool, n)
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	count := 0
	max := 0
	var dfs func(graph [][]int, visited [][]bool, p point)
	dfs = func(graph [][]int, visited [][]bool, p point) {
		if visited[p.x][p.y] || graph[p.x][p.y] == 0 {
			return
		}
		visited[p.x][p.y] = true
		graph[p.x][p.y] = index
		count++
		for _, direction := range directions {
			nx := p.x + direction.x
			ny := p.y + direction.y
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			dfs(graph, visited, point{nx, ny})
		}
	}
	isAllGrid := true // 标记是否整个地图都是陆地
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 0 {
				isAllGrid = false
				continue
			}
			if visited[i][j] {
				continue
			}
			count = 0
			index++
			dfs(graph, visited, point{i, j})
			acr[index] = count
		}
	}
	visitedSet := make(map[int]struct{}, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 || visited[i][j] {
				isAllGrid = false
				continue
			}
			count = 1
			visitedSet = make(map[int]struct{}, 0)
			for _, direction := range directions {
				nx := i + direction.x
				ny := j + direction.y
				if nx < 0 || ny < 0 || nx >= m || ny >= n {
					continue
				}
				if _, ok := visitedSet[graph[nx][ny]]; ok {
					continue
				}
				count += acr[graph[nx][ny]]
				visitedSet[graph[nx][ny]] = struct{}{}
			}
			if count > max {
				max = count
			}
		}
	}
	if isAllGrid {
		fmt.Println(m * n)
	} else {
		fmt.Println(max)
	}

}
