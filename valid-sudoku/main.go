package main

func isValidSudoku(board [][]byte) bool {
	// validate rows
	for i := 0; i < len(board); i++ {
		m := map[byte]interface{}{}
		for j := 0; j < len(board); j++ {
			s := board[i][j]
			if string(s) == "." {
				continue
			}
			s -= '0'
			if s < 1 || s > 9 {
				return false
			}
			if _, ok := m[s]; ok {
				return false
			}
			m[s] = nil
		}
	}
	// validate columns
	for i := 0; i < len(board); i++ {
		m := map[byte]interface{}{}
		for j := 0; j < len(board); j++ {
			s := board[j][i]
			if string(s) == "." {
				continue
			}
			s -= '0'
			if s < 1 || s > 9 {
				return false
			}
			if _, ok := m[s]; ok {
				return false
			}
			m[s] = nil
		}
	}
	// validate boxes
	for i := 0; i < len(board); i = i + 3 {
		for j := 0; j < len(board); j = j + 3 {
			m := map[byte]interface{}{}
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					s := board[k][l]
					if string(s) == "." {
						continue
					}
					s -= '0'
					if s < 1 || s > 9 {
						return false
					}
					if _, ok := m[s]; ok {
						return false
					}
					m[s] = nil
				}
			}
		}
	}
	return true
}
