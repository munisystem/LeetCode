package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			tmp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix)-1-j]
			matrix[i][len(matrix)-1-j] = tmp
		}
	}
}
