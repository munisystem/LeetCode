package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		s := strs[i]
		ss := strings.Split(s, "")
		sort.Strings(ss)
		sorted := strings.Join(ss, "")
		if _, ok := m[sorted]; ok {
			m[sorted] = append(m[sorted], s)
		} else {
			m[sorted] = append([]string{}, s)
		}
	}
	r := make([][]string, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
