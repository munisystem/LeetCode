package main

var symbols map[int]string = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

var units []int = []int{1000, 500, 100, 50, 10, 5, 1}

func intToRoman(num int) string {
	r, t := "", ""
	for i, v := range units {
		if len(t) > 1 {
			r = r + t[0:1]
			t = t[1:]
		}
		cnt := 0
		for num >= v {
			t = t + symbols[v]
			num = num - v
			cnt = cnt + 1
		}
		if cnt == 1 {
			continue
		}
		if cnt < 4 {
			r = r + t
		} else if len(t) == 4 {
			r = r + symbols[v] + symbols[units[i-1]]
		} else {
			r = r + symbols[v] + symbols[units[i-2]]
		}
		t = ""
		cnt = 0
	}
	return r + t
}
