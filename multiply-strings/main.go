package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	a := []int{}
	for i := 0; i < len(num1); i++ {
		s1 := byte(num1[len(num1)-1-i])
		s1 -= '0'
		v1 := int(s1)
		for j := 0; j < len(num2); j++ {
			s2 := byte(num2[len(num2)-1-j])
			s2 -= '0'
			v2 := int(s2)
			digit := i + j + 1
			if digit-1 >= len(a) {
				a = append(a, 0)
			}
			sum := a[digit-1] + (v1 * v2)
			a[digit-1] = sum % 10
			carry := sum / 10
			if carry == 0 {
				continue
			}
			if digit >= len(a) {
				a = append(a, 0)
			}
			a[digit] = a[digit] + carry
		}
	}
	r := ""
	carry := 0
	for i := 0; i < len(a); i++ {
		sum := carry + a[i]
		r = fmt.Sprintf("%d", sum%10) + r
		carry = sum / 10
	}
	if carry > 0 {
		r = fmt.Sprintf("%d", carry) + r
	}
	return r
}
