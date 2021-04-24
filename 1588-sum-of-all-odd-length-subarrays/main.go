package main

func sumOddLengthSubarrays(arr []int) int {
	ans := 0
	for i := 0; i*2+1 <= len(arr); i++ {
		for j := 0; j <= len(arr)-(i*2+1); j++ {
			for k := j; k < j+(i*2+1); k++ {
				ans += arr[k]
			}
		}
	}
	return ans
}
