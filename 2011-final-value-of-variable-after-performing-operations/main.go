package main

func finalValueAfterOperations(operations []string) int {
	var ans int
	for _, operation := range operations {
		if operation[1] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
