package main

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	cols, rows := len(matrix[0]), len(matrix)
	nums := make([]int, cols*rows, cols*rows)
	c, r := 0, 0
	for i := range nums {
		nums[i] = matrix[r][c]
		if (c+r)%2 == 0 {
			if c == cols-1 {
				r++
			} else if r == 0 {
				c++
			} else {
				c++
				r--
			}
		} else {
			if r == rows-1 {
				c++
			} else if c == 0 {
				r++
			} else {
				c--
				r++
			}
		}
	}
	return nums
}
