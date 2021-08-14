package main

func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	ans := make([][]int, cols, cols)
	for i := 0; i < cols; i++ {
		ans[i] = make([]int, rows, rows)
		for j := 0; j < rows; j++ {
			ans[i][j] = matrix[j][i]
		}
	}
	return ans
}
