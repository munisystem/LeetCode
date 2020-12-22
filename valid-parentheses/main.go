package main

var symbols map[string]string = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	q := []string{}
	valid := true
	for i := range s {
		if _, ok := symbols[s[i:i+1]]; ok {
			q = append(q, s[i:i+1])
		} else {
			if len(q) == 0 {
				valid = false
				break
			} else if p := q[len(q)-1]; symbols[p] == s[i:i+1] {
				q = q[0 : len(q)-1]
			} else {
				valid = false
				break
			}
		}
	}
	if len(q) != 0 {
		return false
	}
	return valid
}
