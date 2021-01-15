package main

import (
	"math"
)

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	nums := make([]int, n+1, n+1)
	nums[0], nums[1] = 0, 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
	}
	r := 0
	for i := 0; i <= n; i++ {
		r = int(math.Max(float64(r), float64(nums[i])))
	}
	return r
}
