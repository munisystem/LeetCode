package main

func isIsomorphic(s string, t string) bool {
	m1, m2 := map[int]rune{}, map[int]interface{}{}
	r := make([]rune, len(s), len(s))
	for i := 0; i < len(s); i++ {
		if b, ok := m1[int(s[i])-'A']; ok {
			r[i] = b
		} else if _, ok := m2[int(t[i])-'A']; !ok {
			m1[int(s[i])-'A'] = rune(t[i])
			m2[int(t[i])-'A'] = struct{}{}
			r[i] = m1[int(s[i])-'A']
		} else {
			return false
		}
	}
	return string(r) == t
}
