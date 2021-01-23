package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	m1, m2 := make(map[string]int, 0), make(map[string]int, 0)
	for i := 0; i < len(word1); i++ {
		m1[string(word1[i])]++
		m2[string(word2[i])]++
	}
	if len(m1) != len(m2) {
		return false
	}
	m1c, m2c := make(map[int]int, 0), make(map[int]int, 0)
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
		m1c[v]++
		m2c[m2[k]]++
	}
	for k, v := range m1c {
		m2cv, ok := m2c[k]
		if !ok {
			return false
		} else if v != m2cv {
			return false
		}
	}
	return true
}
