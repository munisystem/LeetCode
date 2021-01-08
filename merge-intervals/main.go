package main

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	r := [][]int{}
	s := intervals[0][0]
	e := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if start <= e {
			e = int(math.Max(float64(e), float64(end)))
			continue
		}
		r = append(r, append([]int{}, s, e))
		s = start
		e = end
	}
	r = append(r, append([]int{}, s, e))
	return r
}
