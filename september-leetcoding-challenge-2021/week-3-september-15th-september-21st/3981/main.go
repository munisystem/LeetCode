package main

func tictactoe(moves [][]int) string {
	aRow := make([]int, 3, 3)
	bRow := make([]int, 3, 3)
	aCol := make([]int, 3, 3)
	bCol := make([]int, 3, 3)
	aD1, aD2 := 0, 0
	bD1, bD2 := 0, 0
	for i := 0; i < len(moves); i++ {
		r, c := moves[i][0], moves[i][1]
		if i%2 == 0 {
			aRow[r]++
			aCol[c]++
			if r == c {
				aD1++
			}
			if r+c == 2 {
				aD2++
			}
			if aRow[r] == 3 || aCol[c] == 3 || aD1 == 3 || aD2 == 3 {
				return "A"
			}
		} else {
			bRow[r]++
			bCol[c]++
			if r == c {
				bD1++
			}
			if r+c == 2 {
				bD2++
			}
			if bRow[r] == 3 || bCol[c] == 3 || bD1 == 3 || bD2 == 3 {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
