package main

func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	r := 0
	for i := 1; i <= n; i++ {
		r += numTrees(i-1) * numTrees(n-i)
	}
	return r
}
