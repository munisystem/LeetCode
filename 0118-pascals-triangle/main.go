package main

import "fmt"

func main() {
	fmt.Println(generate(6))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}
