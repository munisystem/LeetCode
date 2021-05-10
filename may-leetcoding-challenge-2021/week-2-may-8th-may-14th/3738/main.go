package main

func countPrimes(n int) int {
	nums := make([]bool, n, n)
	for i := 2; i*i < n; i++ {
		if nums[i] == true {
			continue
		}
		for j := 2; i*j < n; j++ {
			nums[i*j] = true
		}
	}
	cnt := 0
	for i := 2; i < len(nums); i++ {
		if nums[i] == false {
			cnt++
		}
	}
	return cnt
}
