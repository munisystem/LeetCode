package main

func reverseWords(s string) string {
	r := ""
	t := ""
	for i := len(s) - 1; i >= 0; i-- {
		if t == "" && string(s[i]) == " " {
			continue
		} else if string(s[i]) == " " {
			r = join(r, t)
			t = ""
		} else {
			t = string(s[i]) + t
		}
	}
	if t != "" {
		r = join(r, t)
	}
	return r
}

func join(s, w string) string {
	if len(s) == 0 {
		return s + w
	} else {
		return s + " " + w
	}
}
