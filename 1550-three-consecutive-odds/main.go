package main

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			count = 0
			continue
		}
		count++
		if count >= 3 {
			return true
		}
	}
	return false
}
