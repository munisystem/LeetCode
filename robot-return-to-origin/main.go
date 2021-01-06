package main

func judgeCircle(moves string) bool {
	m := make(map[string]int, 4)
	m["U"], m["D"], m["L"], m["R"] = 0, 0, 0, 0
	for i := 0; i < len(moves); i++ {
		m[string(moves[i])]++
	}
	if m["U"] == m["D"] && m["L"] == m["R"] {
		return true
	}
	return false
}
