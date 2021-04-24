
var parentheses = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func isValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "(" || char == "[" || char == "{" {
			stack = append(stack, char)
			continue
		}
		if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != parentheses[char] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
