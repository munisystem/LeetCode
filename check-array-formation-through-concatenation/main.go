package main

func canFormArray(arr []int, pieces [][]int) bool {
	for i := 0; i < len(pieces); i++ {
		start := -1
		for j := 0; j < len(arr); j++ {
			if arr[j] == pieces[i][0] {
				start = j
				break
			}
		}
		if start == -1 {
			return false
		}
		end := start + len(pieces[i]) - 1
		if end >= len(arr) {
			return false
		}
		for j := 0; j < len(pieces[i]); j++ {
			if pieces[i][j] != arr[start+j] {
				return false
			}
		}
		arr = append(arr[:start], arr[end:]...)
	}
	return true
}
