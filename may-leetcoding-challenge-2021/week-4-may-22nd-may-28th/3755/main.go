package main

import "strconv"

var operators = map[string]interface{}{
	"+": struct{}{},
	"-": struct{}{},
	"*": struct{}{},
	"/": struct{}{},
}

func evalRPN(tokens []string) int {
	queue := make([]int, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		_, ok := operators[tokens[i]]
		if !ok {
			num, _ := strconv.ParseInt(tokens[i], 10, 0)
			queue = append(queue, int(num))
			continue
		}
		a, b := queue[len(queue)-2], queue[len(queue)-1]
		queue = queue[:len(queue)-2]
		switch tokens[i] {
		case "+":
			queue = append(queue, a+b)
		case "-":
			queue = append(queue, a-b)
		case "*":
			queue = append(queue, a*b)
		case "/":
			queue = append(queue, a/b)
		}
	}
	return queue[0]
}
