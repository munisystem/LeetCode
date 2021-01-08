package main

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	m := make(map[int]int, 0)
	starts := []int{}
	for i := 0; i < len(intervals); i++ {
		s, e := intervals[i][0], intervals[i][1]
		if v, ok := m[s]; ok {
			m[s] = int(math.Max(float64(v), float64(e)))
		} else {
			m[s] = e
			starts = append(starts, s)
		}
	}
	sort.Ints(starts)

	r := [][]int{}
	s := starts[0]
	e := m[s]
	for i := 1; i < len(starts); i++ {
		if ss := starts[i]; ss <= e {
			e = int(math.Max(float64(e), float64(m[ss])))
			continue
		}
		r = append(r, append([]int{}, s, e))
		s = starts[i]
		e = m[s]
	}
	r = append(r, append([]int{}, s, e))
	return r
}
