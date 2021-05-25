package main

import "strconv"

func evalRPN(tokens []string) int {
	queue := make([]int, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		var a, b int
		switch t := tokens[i]; t {
		case "+":
			a, b, queue = pop(queue)
			queue = append(queue, a+b)
		case "-":
			a, b, queue = pop(queue)
			queue = append(queue, a-b)
		case "*":
			a, b, queue = pop(queue)
			queue = append(queue, a*b)
		case "/":
			a, b, queue = pop(queue)
			queue = append(queue, a/b)
		default:
			num, _ := strconv.ParseInt(t, 10, 0)
			queue = append(queue, int(num))
		}
	}
	return queue[0]
}

func pop(queue []int) (int, int, []int) {
	a, b := queue[len(queue)-2], queue[len(queue)-1]
	return a, b, queue[:len(queue)-2]
}
