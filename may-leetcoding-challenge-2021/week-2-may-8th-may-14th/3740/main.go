package main

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	return NumMatrix{matrix: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.matrix[i][col2]
		if col1 > 0 {
			sum -= this.matrix[i][col1-1]
		}
	}
	return sum
}
