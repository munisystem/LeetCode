package main

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		n := int(num[i]) - '0'
		if n%2 == 1 {
			return num[0 : i+1]
		}
	}
	return ""
}
