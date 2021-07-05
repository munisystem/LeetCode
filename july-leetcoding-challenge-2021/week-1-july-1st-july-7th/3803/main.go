package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows*cols != r*c {
		return mat
	}
	m := make([][]int, r, r)
	for i := 0; i < r; i++ {
		m[i] = make([]int, c, c)
	}
	for i := 0; i < r*c; i++ {
		m[i/c][i%c] = mat[i/cols][i%cols]
	}
	return m
}
