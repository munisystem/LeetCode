package main

import "strings"

func commonChars(A []string) []string {
	if len(A) == 1 {
		return strings.Split(A[0], "")
	}
	mm := make([]map[int]int, len(A)-1, len(A)-1)
	for i := 1; i < len(A); i++ {
		m := map[int]int{}
		for j := 0; j < len(A[i]); j++ {
			m[int(A[i][j])-'a']++
		}
		mm[i-1] = m
	}
	ans := []string{}
	for i := 0; i < len(A[0]); i++ {
		isExist := true
		for j := 0; j < len(mm); j++ {
			if mm[j][int(A[0][i])-'a'] > 0 {
				mm[j][int(A[0][i])-'a']--
			} else {
				isExist = false
				break
			}
		}
		if isExist {
			ans = append(ans, string(A[0][i]))
		}
	}
	return ans
}
