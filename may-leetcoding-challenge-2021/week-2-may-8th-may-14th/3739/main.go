package main

import "math"

func maxScore(cardPoints []int, k int) int {
	ans, total, sum := 0, 0, 0
	for i := 0; i < len(cardPoints); i++ {
		total += cardPoints[i]
	}
	if len(cardPoints) == k {
		return total
	}
	for i := 0; i < len(cardPoints); i++ {
		sum += cardPoints[i]
		if i < len(cardPoints)-k-1 {
			continue
		}
		ans = int(math.Max(float64(ans), float64(total-sum)))
		sum -= cardPoints[i-(len(cardPoints)-k-1)]
	}
	return ans
}
