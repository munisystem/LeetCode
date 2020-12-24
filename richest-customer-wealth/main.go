package main

func maximumWealth(accounts [][]int) int {
	return max(accounts)
}

func max(accounts [][]int) int {
	if len(accounts) <= 1 {
		return sum(accounts[0])
	}
	i := (len(accounts) - 1) / 2
	h, d := max(accounts[0:i+1]), max(accounts[i+1:len(accounts)])
	if h >= d {
		return h
	}
	return d
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s = s + n
	}
	return s
}
