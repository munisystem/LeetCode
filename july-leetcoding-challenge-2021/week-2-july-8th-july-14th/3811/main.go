package main

func isIsomorphic(s string, t string) bool {
	m1, m2 := map[int]byte{}, map[int]struct{}{}
	str := make([]byte, len(s), len(s))
	for i := 0; i < len(s); i++ {
		if v, ok := m1[int(s[i])-'A']; ok {
			str[i] = v
		} else if _, ok := m2[int(t[i])-'A']; !ok {
			str[i] = t[i]
			m1[int(s[i])-'A'] = t[i]
			m2[int(t[i])-'A'] = struct{}{}
		} else {
			return false
		}
	}
	return string(str) == t
}
