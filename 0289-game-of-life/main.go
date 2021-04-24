package main

// Constraints:
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 25
// board[i][j] is 0 or 1.

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	r := make([][]int, m, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			r[i][j] = board[i][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lives := 0
			if j > 0 {
				lives = lives + r[i][j-1]
			}
			if j < n-1 {
				lives = lives + r[i][j+1]
			}
			if i > 0 {
				lives = lives + r[i-1][j]
				if j > 0 {
					lives = lives + r[i-1][j-1]
				}
				if j < n-1 {
					lives = lives + r[i-1][j+1]
				}
			}
			if i < m-1 {
				lives = lives + r[i+1][j]
				if j > 0 {
					lives = lives + r[i+1][j-1]
				}
				if j < n-1 {
					lives = lives + r[i+1][j+1]
				}
			}
			l := r[i][j]
			switch {
			case l == 0 && lives == 3:
				board[i][j] = 1
			case l == 1 && (lives == 2 || lives == 3):
				board[i][j] = 1
			default:
				board[i][j] = 0
			}
		}
	}
}
