package main

import "math"

func minPartitions(n string) int {
	r := 0
	for i := 0; i < len(n); i++ {
		v := int(byte(n[i]) - '0')
		r = int(math.Max(float64(r), float64(v)))
	}
	return r
}
