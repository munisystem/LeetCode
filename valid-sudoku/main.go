package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]interface{}, 9, 9)
	columns := make([]map[byte]interface{}, 9, 9)
	boxes := make([]map[byte]interface{}, 9, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]interface{}, 0)
		columns[i] = make(map[byte]interface{}, 0)
		boxes[i] = make(map[byte]interface{}, 0)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; i < len(board); i++ {
			s := board[i][j]
			if string(s) == "." {
				continue
			}
			s -= '0'
			if s < 1 || s > 9 {
				return false
			}
			bi := (i/3)*3 + (j / 3)
			if _, ok := rows[j][s]; ok {
				return false
			} else {
				rows[j][s] = nil
			}
			if _, ok := columns[i][s]; ok {
				return false
			} else {
				columns[i][s] = nil
			}
			if _, ok := boxes[bi][s]; ok {
				return false
			} else {
				boxes[bi][s] = nil
			}
		}
	}
	return true
}
