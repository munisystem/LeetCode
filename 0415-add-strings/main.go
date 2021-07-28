package main

func addStrings(num1 string, num2 string) string {
	var nums []byte
	var carry int
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i]) - '0'
		}
		if j >= 0 {
			n2 = int(num2[j]) - '0'
		}
		n := (carry + n1 + n2)
		nums = append(nums, byte(n%10+'0'))
		carry = n / 10
		i--
		j--
	}
	if carry > 0 {
		nums = append(nums, byte(carry+'0'))
	}
	ans := make([]byte, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		ans[len(nums)-1-i] = nums[i]
	}
	return string(ans)
}
