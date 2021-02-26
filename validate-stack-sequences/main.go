package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	pi := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[pi] {
			stack = stack[:len(stack)-1]
			pi++
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
