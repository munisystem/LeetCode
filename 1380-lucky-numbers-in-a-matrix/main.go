package main

func luckyNumbers(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	ans := make([]int, 0)
	for i := 0; i < rows; i++ {
		minv := matrix[i][0]
		minc := 0
		for j := 1; j < cols; j++ {
			if minv > matrix[i][j] {
				minv = matrix[i][j]
				minc = j
			}
		}
		for j := 0; j < rows; j++ {
			if matrix[j][minc] > minv {
				minv = 0
				break
			}
		}
		if minv != 0 {
			ans = append(ans, minv)
		}
	}
	return ans
}
