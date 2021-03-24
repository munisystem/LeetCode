package main

import "sort"

func advantageCount(A []int, B []int) []int {
	m := map[int][]int{}
	for i := 0; i < len(B); i++ {
		m[B[i]] = append(m[B[i]], i)
	}
	sort.Ints(A)
	sort.Ints(B)
	ans := make([]int, len(A), len(A))
	l, r := 0, len(A)-1
	for i := 0; i < len(A); i++ {
		var bi int
		if A[i] > B[l] {
			bi = l
			l++
		} else {
			bi = r
			r--
		}
		j := m[B[bi]][0]
		m[B[bi]] = m[B[bi]][1:]
		ans[j] = A[i]
	}
	return ans
}
