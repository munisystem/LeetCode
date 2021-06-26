package main

func titleToNumber(columnTitle string) int {
	var ans int
	for i := 0; i < len(columnTitle); i++ {
		ans = ans*26 + (int(columnTitle[i]) - 'A' + 1)
	}
	return ans
}
