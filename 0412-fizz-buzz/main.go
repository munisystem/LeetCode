package main

import "fmt"

func fizzBuzz(n int) []string {
	ans := make([]string, n, n)
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if len(s) == 0 {
			s = fmt.Sprintf("%d", i)
		}
		ans[i-1] = s
	}
	return ans
}
