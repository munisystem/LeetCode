package main

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	max := 0
	i, j := 0, len(height)-1
	for {
		if i == j {
			break
		}
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i = i + 1
		} else {
			area = height[j] * (j - i)
			j = j - 1
		}
		if area > max {
			max = area
		}
	}
	return max
}
