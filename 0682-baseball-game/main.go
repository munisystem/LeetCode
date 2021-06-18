package main

import "strconv"

func calPoints(ops []string) int {
	queue := []int{}
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			queue = queue[:len(queue)-1]
		case "D":
			queue = append(queue, queue[len(queue)-1]*2)
		case "+":
			queue = append(queue, queue[len(queue)-1]+queue[len(queue)-2])
		default:
			num, _ := strconv.Atoi(ops[i])
			queue = append(queue, num)
		}
	}
	var ans int
	for i := 0; i < len(queue); i++ {
		ans += queue[i]
	}
	return ans
}
