package main

import "fmt"

func main() {
	fmt.Println(generate(6))
}

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		r := []int{}
		for j := 0; j <= i; j++ {
			s := 1
			for k := 0; k < j; k++ {
				s = s * (i - k)
				s = s / (k + 1)
			}
			r = append(r, s)
		}
		res = append(res, r)
	}
	return res
}
