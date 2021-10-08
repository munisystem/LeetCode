package main

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	ans := make([][]int, m, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n, n)
	}
	for i := 0; i < len(original); i++ {
		ans[i/n][i%n] = original[i]
	}
	return ans
}
