package main

func judgeCircle(moves string) bool {
	t := make([]int, 26, 26)
	for i := 0; i < len(moves); i++ {
		t[moves[i]-'A']++
	}
	if t['U'-'A'] == t['D'-'A'] && t['L'-'A'] == t['R'-'A'] {
		return true
	}
	return false
}
