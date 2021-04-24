package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				m := make([]map[int]bool, len(board), len(board))
				for i := 0; i < len(board); i++ {
					m[i] = make(map[int]bool, len(board[0]))
					for j := 0; j < len(board[0]); j++ {
						m[i][j] = false
					}
				}
				m[i][j] = true
				b := backtrack(board, word[1:], i, j, m)
				if b {
					return true
				}
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, col, row int, m []map[int]bool) bool {
	if len(word) == 0 {
		return true
	}
	if col < len(board)-1 && !m[col+1][row] && board[col+1][row] == word[0] {
		cm := make([]map[int]bool, len(board), len(board))
		for i := 0; i < len(board); i++ {
			cm[i] = make(map[int]bool, len(board[0]))
			for j := 0; j < len(board[0]); j++ {
				cm[i][j] = m[i][j]
			}
		}
		cm[col+1][row] = true
		b := backtrack(board, word[1:], col+1, row, cm)
		if b {
			return true
		}
	}
	if col > 0 && !m[col-1][row] && board[col-1][row] == word[0] {
		cm := make([]map[int]bool, len(board), len(board))
		for i := 0; i < len(board); i++ {
			cm[i] = make(map[int]bool, len(board[0]))
			for j := 0; j < len(board[0]); j++ {
				cm[i][j] = m[i][j]
			}
		}
		cm[col-1][row] = true
		b := backtrack(board, word[1:], col-1, row, cm)
		if b {
			return true
		}
	}
	if row < len(board[0])-1 && !m[col][row+1] && board[col][row+1] == word[0] {
		cm := make([]map[int]bool, len(board), len(board))
		for i := 0; i < len(board); i++ {
			cm[i] = make(map[int]bool, len(board[0]))
			for j := 0; j < len(board[0]); j++ {
				cm[i][j] = m[i][j]
			}
		}
		cm[col][row+1] = true
		b := backtrack(board, word[1:], col, row+1, cm)
		if b {
			return true
		}
	}
	if row > 0 && !m[col][row-1] && board[col][row-1] == word[0] {
		cm := make([]map[int]bool, len(board), len(board))
		for i := 0; i < len(board); i++ {
			cm[i] = make(map[int]bool, len(board[0]))
			for j := 0; j < len(board[0]); j++ {
				cm[i][j] = m[i][j]
			}
		}
		cm[col][row-1] = true
		b := backtrack(board, word[1:], col, row-1, cm)
		if b {
			return true
		}
	}
	return false
}
