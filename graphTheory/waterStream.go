package main

import "fmt"

func waterStream() {
	var (
		m int
		n int
	)
	type point struct {
		x int
		y int
	}
	directions := []point{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	fmt.Scanln(&m, &n)
	graph := make([][]int, m)
	canReach1 := make([][]bool, m)
	canReach2 := make([][]bool, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
		canReach1[i] = make([]bool, n)
		canReach2[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	var dfs func(graph [][]int, visited [][]bool, x, y int)

	dfs = func(graph [][]int, visited [][]bool, x, y int) {
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		for _, direction := range directions {
			nx := x + direction.x
			ny := y + direction.y
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			if graph[nx][ny] < graph[x][y] {
				continue
			}
			dfs(graph, visited, nx, ny)
		}
		return
	}
	for i := 0; i < m; i++ {
		dfs(graph, canReach1, i, 0)
		dfs(graph, canReach2, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(graph, canReach1, 0, i)
		dfs(graph, canReach2, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReach1[i][j] && canReach2[i][j] {
				fmt.Printf("%d %d\n", i, j)
			}
		}
	}
}
