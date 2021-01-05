package main

func defangIPaddr(address string) string {
	s := ""
	for i := 0; i < len(address); i++ {
		if string(address[i]) == "." {
			s = s + "[.]"
		} else {
			s = s + string(address[i])
		}
	}
	return s
}
