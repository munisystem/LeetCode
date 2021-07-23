package main

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		n := i
		isDivisible := true
		for n > 0 {
			if n%10 == 0 || i%(n%10) != 0 {
				isDivisible = false
				break
			}
			n /= 10
		}
		if isDivisible {
			ans = append(ans, i)
		}
	}
	return ans
}
