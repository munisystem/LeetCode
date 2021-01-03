package main

import "math"

// MEMO: I must not use the Dynamic Programming.
// > Follow up:
// >   If you have figured out the O(n) solution,
// >   try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	r := math.MinInt32
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			s = s + nums[j]
			r = int(math.Max(float64(r), float64(s)))
		}
	}
	if r >= math.MaxInt32 {
		return math.MaxInt32
	} else if r <= math.MinInt32 {
		return math.MinInt32
	}
	return r
}
