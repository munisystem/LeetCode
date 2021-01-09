package main

import (
	"math"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	r := [][]int{}
	s, e := newInterval[0], newInterval[1]
	i := 0
	for ; i < len(intervals); i++ {
		is, ie := intervals[i][0], intervals[i][1]
		if ie < s {
			r = append(r, append([]int{}, is, ie))
			continue
		}
		if is > e {
			r = append(r, append([]int{}, s, e))
			s, e = is, ie
			continue
		}
		if is <= e {
			s = int(math.Min(float64(s), float64(is)))
			e = int(math.Max(float64(e), float64(ie)))
			continue
		}
		r = append(r, append([]int{}, s, e))
		s, e = is, ie
	}
	r = append(r, append([]int{}, s, e))
	return r
}
