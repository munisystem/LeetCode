package main

func countBalls(lowLimit int, highLimit int) int {
	boxes := map[int]int{}
	max := 0
	for i := lowLimit; i <= highLimit; i++ {
		number := i
		box := 0
		for number > 0 {
			box += number % 10
			number = number / 10
		}
		boxes[box]++
		if boxes[box] > max {
			max = boxes[box]
		}
	}
	return max
}
