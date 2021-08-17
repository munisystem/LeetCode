package main

func reformatNumber(number string) string {
	nums := []rune{}
	for _, n := range number {
		switch i := n - '0'; {
		case i >= 0 && i <= 9:
			nums = append(nums, n)
		default:
			continue
		}
	}

	if len(nums) < 4 {
		return string(nums)
	} else if len(nums) == 4 {
		return string(nums[:2]) + "-" + string(nums[2:])
	}

	var ans []rune
	i := 0
	for i < len(nums) {
		ans = append(ans, nums[i:i+3]...)
		ans = append(ans, '-')
		i += 3

		left := len(nums[i:])
		if left <= 4 {
			if left <= 3 {
				ans = append(ans, nums[i:]...)
			} else {
				ans = append(ans, nums[i:i+2]...)
				ans = append(ans, '-')
				ans = append(ans, nums[i+2:]...)
			}
			break
		}
	}
	return string(ans)
}
