package main

func majorityElement(nums []int) int {
	// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
	m, c := 0, 0
	for i := 0; i < len(nums); i++ {
		if c == 0 {
			m, c = nums[i], 1
		} else if m == nums[i] {
			c++
		} else {
			c--
		}
	}
	return m
}
