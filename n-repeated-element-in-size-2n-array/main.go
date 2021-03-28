package main

func repeatedNTimes(A []int) int {
	m := map[int]int{}
	var ans int
	for i := 0; i < len(A); i++ {
		m[A[i]]++
		if m[A[i]] > 1 {
			ans = A[i]
			break
		}
	}
	return ans
}
