package main

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	r := ""
	carry := 0
	for i >= 0 || j >= 0 {
		ai, bi := 0, 0
		if i >= 0 && string(a[i]) == "1" {
			ai = 1
		}
		if j >= 0 && string(b[j]) == "1" {
			bi = 1
		}
		if (ai+bi+carry)/2 != 0 {
			if (ai+bi+carry)%2 == 0 {
				r = "0" + r
			} else {
				r = "1" + r
			}
			carry = 1
		} else {
			if (ai+bi+carry)%2 == 0 {
				r = "0" + r
			} else {
				r = "1" + r
			}
			carry = 0
		}
		i--
		j--
	}
	if carry > 0 {
		r = "1" + r
	}
	return r
}
