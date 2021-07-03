package main

import "fmt"

func convertToBase7(num int) string {
	var ans string
	if num == 0 {
		return "0"
	}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	for num > 0 {
		ans = fmt.Sprintf("%d%s", num%7, ans)
		num /= 7
	}
	if isNegative {
		ans = "-" + ans
	}
	return ans
}
