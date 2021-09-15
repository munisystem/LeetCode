package main

import "math"

func maxTurbulenceSize(arr []int) int {
	var ans int
	ans, inc, dec := 1, 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			inc = dec + 1
			dec = 1
		} else if arr[i-1] > arr[i] {
			dec = inc + 1
			inc = 1
		} else {
			inc, dec = 1, 1
		}
		ans = int(math.Max(float64(ans), math.Max(float64(inc), float64(dec))))
	}
	return ans
}
