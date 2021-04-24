package main

var symbols map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	r, st, sz, l := 0, 0, 0, 0
	for i := range s {
		c := symbols[s[i:i+1]]
		if l == 0 {
			st, l, sz = c, c, sz+1
			continue
		}
		if l == c {
			st, sz = st+c, sz+1
			continue
		}
		if l > c {
			r = r + st
			st, l, sz = c, c, 1
			continue
		}
		if sz == 1 {
			r = r + c - st
			st, l, sz = 0, 0, 0
			continue
		} else {
			r = r + st
			st, l, sz = c, c, 1
			continue
		}
	}
	return r + st
}
