package main

func solveSudo(board [][]byte) {
	len := 9
	isValid := func(row, col int, num byte, board [][]byte) bool {
		for i := 0; i < len; i++ {
			if board[i][col] == num {
				return false
			}
		}
		for i := 0; i < len; i++ {
			if board[row][i] == num {
				return false
			}
		}
		startRow := row - row%3
		startCol := col - col%3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == num {
					return false
				}
			}
		}
		return true
	}
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := 0; i < len; i++ {
			for j := 0; j < len; j++ {
				if board[i][j] != '.' {
					continue
				}

				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backtracking(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking(board)

}
