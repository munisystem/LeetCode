package main

func isPossible(target []int) bool {
	sum, max := 0, 0
	for i := 0; i < len(target); i++ {
		if target[i] > target[max] {
			max = i
		}
		sum += target[i]
	}
	for sum > 1 && target[max] > sum/2 {
		sum -= target[max]
		if sum <= 1 {
			return sum == 1
		}
		target[max] = target[max] % sum
		sum += target[max]
		for i := 0; i < len(target); i++ {
			if target[i] > target[max] {
				max = i
			}
		}
	}
	return sum == len(target)
}
