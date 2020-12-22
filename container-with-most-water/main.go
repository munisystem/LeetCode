package main

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	max := 0
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			m := 0
			if height[i] < height[j] {
				m = height[i]
			} else {
				m = height[j]
			}
			area := int(m) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}
