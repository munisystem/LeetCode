package main

func addStrings(num1 string, num2 string) string {
	var ans string
	var carry int
	i1, i2 := len(num1)-1, len(num2)-1
	for i1 >= 0 || i2 >= 0 {
		var n1, n2 int
		if i1 >= 0 {
			n1 = int(num1[i1]) - '0'
			i1--
		}
		if i2 >= 0 {
			n2 = int(num2[i2]) - '0'
			i2--
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		ans = string(sum%10+'0') + ans
	}
	if carry > 0 {
		ans = string(carry+'0') + ans
	}
	return ans
}
