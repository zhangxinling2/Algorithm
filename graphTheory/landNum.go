package main

import "fmt"

var (
	dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func landNum1() {
	var (
		n int
		m int
	)
	fmt.Scanln(&n, &m)
	graph := make([][]int, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		visited[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	var dfs func(graph [][]int, visited [][]bool, x, y int)
	dfs = func(graph [][]int, visited [][]bool, x, y int) {
		for i := 0; i < 4; i++ {
			nextX := x + dir[i][0]
			nextY := y + dir[i][1]
			if nextX < 0 || nextX >= len(graph) || nextY < 0 || nextY >= len(graph[0]) {
				continue
			}
			if !visited[nextX][nextY] && graph[nextX][nextY] == 1 {
				visited[nextX][nextY] = true
				dfs(graph, visited, nextX, nextY)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && graph[i][j] == 1 {
				res++
				dfs(graph, visited, i, j)
			}
		}
	}
	fmt.Println(res)
}
func landNum2() {
	var (
		n int
		m int
	)
	fmt.Scanln(&n, &m)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	type point struct {
		x, y int
	}
	var directions = []point{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	var visited = map[point]bool{}
	// bfs 需要数据结构来保存遍历顺序，一般使用队列
	var queue = make([]point, 0)
	var bfs func(graph [][]int, visited map[point]bool, queue []point)
	var ans int
	bfs = func(graph [][]int, visited map[point]bool, queue []point) {
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, u := range directions {
				nextX := p.x + u.x
				nextY := p.y + u.y
				if nextX < 0 || nextX >= n || nextY < 0 || nextY >= m {
					continue
				}
				if !visited[point{nextX, nextY}] && graph[nextX][nextY] == 1 {
					visited[point{nextX, nextY}] = true
					queue = append(queue, point{nextX, nextY})
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 1 && !visited[point{i, j}] {
				visited[point{i, j}] = true
				queue = append(queue, point{i, j})
				ans++
				bfs(graph, visited, queue)
			}
		}
	}
	fmt.Println(ans)
}
