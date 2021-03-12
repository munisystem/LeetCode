package main

func interpret(command string) string {
	ans := ""
	t := ""
	for i := 0; i < len(command); i++ {
		switch string(command[i]) {
		case "G":
			ans += "G"
		case ")":
			if len(t) == 1 {
				ans += "o"
			} else {
				ans += "al"
			}
			t = ""
		default:
			t += string(command[i])
		}
	}
	return ans
}
