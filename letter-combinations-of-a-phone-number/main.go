package main

var m = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	r := []string{""}
	for i := 0; i < len(digits); i++ {
		sub := []string{}
		for j := 0; j < len(m[string(digits[i])]); j++ {
			for k := 0; k < len(r); k++ {
				sub = append(sub, r[k]+m[string(digits[i])][j])
			}
		}
		r = append([]string{}, sub...)
	}
	return r
}
