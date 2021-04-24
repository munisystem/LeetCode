package main

func diagonalSum(mat [][]int) int {
	ans := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if j == i || j == len(mat)-1-i {
				ans += mat[i][j]
			}
		}
	}
	return ans
}
