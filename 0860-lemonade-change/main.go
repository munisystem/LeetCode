package main

func lemonadeChange(bills []int) bool {
	var five, ten int
	for i := 0; i < len(bills); i++ {
		bill := bills[i]
		switch bill {
		case 10:
			if five >= 1 {
				five--
				ten++
			} else {
				return false
			}
		case 20:
			if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		default:
			five++
		}
	}
	return true
}
