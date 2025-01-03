package main

import "fmt"

func landAcreage() {
	type point struct {
		x int
		y int
	}
	var (
		m int
		n int
	)
	var directions = []point{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	acr := 0
	max := 0
	fmt.Scanln(&m, &n)
	graph := make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var bfs func(graph [][]int, visited [][]bool, x, y int)
	bfs = func(graph [][]int, visited [][]bool, x, y int) {
		for _, direction := range directions {
			tx := x + direction.x
			ty := y + direction.y
			if tx < 0 || ty < 0 || tx >= len(graph) || ty >= len(graph[0]) {
				continue
			}
			if graph[tx][ty] == 1 && visited[tx][ty] == false {
				acr++
				visited[tx][ty] = true
				bfs(graph, visited, tx, ty)
				if acr > max {
					max = acr
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 && visited[i][j] == false {
				acr = 0
				bfs(graph, visited, i, j)
			}
		}
	}
	fmt.Println(max)
}

func landAcreageII() {
	type point struct {
		x int
		y int
	}
	var (
		m int
		n int
	)
	var directions = []point{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	acr := 0
	max := 0
	fmt.Scanln(&m, &n)
	graph := make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	queue := make([]point, 0)
	var dfs func(graph [][]int, visited [][]bool, queue []point)
	dfs = func(graph [][]int, visited [][]bool, queue []point) {
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				tx := p.x + direction.x
				ty := p.y + direction.y
				if tx < 0 || ty < 0 || tx >= len(graph) || ty >= len(graph[0]) {
					continue
				}
				if graph[tx][ty] == 1 && visited[tx][ty] == false {
					acr++
					if acr > max {
						max = acr
					}
					visited[tx][ty] = true
					queue = append(queue, point{tx, ty})
					dfs(graph, visited, queue)

				}
			}
		}

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 && visited[i][j] == false {
				acr = 0
				queue = append(queue, point{i, j})
				dfs(graph, visited, queue)
			}
		}
	}
	fmt.Println(max)
}
