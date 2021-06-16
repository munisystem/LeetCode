package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	ans := make([][]int, r, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c, c)
	}
	for i := 0; i < r*c; i++ {
		ans[i/c][i%c] = mat[i/m][i%m]
	}
	return ans
}
