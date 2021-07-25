package main

func decrypt(code []int, k int) []int {
	ans := make([]int, len(code), len(code))
	start, end := 1, k
	if k < 0 {
		k *= -1
		start = len(code) - k
		end = len(code) - 1
	}
	var sum int
	for i := start; i <= end; i++ {
		sum += code[i]
	}
	for i := 0; i < len(code); i++ {
		ans[i] = sum
		sum -= code[start%len(code)]
		start++
		end++
		sum += code[end%len(code)]
	}
	return ans
}
