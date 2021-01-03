package main

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		s := 1
		for j := 1; j <= i; j++ {
			s = s * (rowIndex + 1 - j)
			s = s / j
		}
		res[i] = s
	}
	return res
}
