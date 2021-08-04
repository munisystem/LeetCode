package main

func convertToTitle(columnNumber int) string {
	var ans string
	for columnNumber > 0 {
		if columnNumber%26 == 0 {
			ans = "Z" + ans
			columnNumber /= 26
			columnNumber--
		} else {
			ans = string(columnNumber%26+'A'-1) + ans
			columnNumber /= 26
		}
	}
	return ans
}
