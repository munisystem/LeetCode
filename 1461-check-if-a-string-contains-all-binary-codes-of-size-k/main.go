package main

func hasAllCodes(s string, k int) bool {
	n := 1 << k
	m := map[string]interface{}{}
	for i := 0; i <= len(s)-k; i++ {
		sub := s[i : i+k]
		if _, ok := m[sub]; !ok {
			m[sub] = struct{}{}
			n--
			if n == 0 {
				return true
			}
		}
	}
	return false
}
