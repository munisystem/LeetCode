package main

func countNegatives(grid [][]int) int {
	var ans int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				ans++
			}
		}
	}
	return ans
}
