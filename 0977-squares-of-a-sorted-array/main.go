package main

func sortedSquares(nums []int) []int {
	squares := []int{}
	for i := 0; i < len(nums); i++ {
		s := nums[i] * nums[i]
		if len(squares) == 0 {
			squares = append(squares, s)
			continue
		}
		if s >= squares[len(squares)-1] {
			squares = append(squares, s)
			continue
		}
		for j := 0; j < len(squares); j++ {
			if s <= squares[j] {
				ns := []int{}
				head, tail := squares[0:j], squares[j:]
				ns = append(ns, head...)
				ns = append(ns, s)
				ns = append(ns, tail...)
				squares = ns
				break
			}
		}
	}
	return squares
}
