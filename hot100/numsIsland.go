package main

func numIslands(grid [][]byte) int {
	cnt := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		if grid[i][j] == '0' {
			return
		}
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				if grid[i][j] == '1' {
					cnt++
				}
				dfs(grid, i, j)
			} else {
				continue
			}
		}
	}
	return cnt
}
