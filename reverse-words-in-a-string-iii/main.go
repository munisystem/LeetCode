package main

func reverseWords(s string) string {
	ans := []byte{}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			for j := len(stack) - 1; j >= 0; j-- {
				ans = append(ans, stack[j])
			}
			ans = append(ans, ' ')
			stack = []byte{}
			continue
		}
		stack = append(stack, s[i])
	}
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return string(ans)
}
