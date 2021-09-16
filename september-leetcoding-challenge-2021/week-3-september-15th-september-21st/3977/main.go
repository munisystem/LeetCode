package main

func spiralOrder(matrix [][]int) []int {
	is, ie, js, je := 0, len(matrix)-1, 0, len(matrix[0])-1
	r := []int{}
	for {
		for i, j := is, js; j <= je; j++ {
			r = append(r, matrix[i][j])
			if len(r) == len(matrix)*len(matrix[0]) {
				return r
			}
		}
		for i, j := is+1, je; i <= ie; i++ {
			r = append(r, matrix[i][j])
			if len(r) == len(matrix)*len(matrix[0]) {
				return r
			}
		}
		for i, j := ie, je-1; j >= js; j-- {
			r = append(r, matrix[i][j])
			if len(r) == len(matrix)*len(matrix[0]) {
				return r
			}
		}
		for i, j := ie-1, js; i >= is+1; i-- {
			r = append(r, matrix[i][j])
			if len(r) == len(matrix)*len(matrix[0]) {
				return r
			}
		}
		is++
		ie--
		js++
		je--
	}
}
