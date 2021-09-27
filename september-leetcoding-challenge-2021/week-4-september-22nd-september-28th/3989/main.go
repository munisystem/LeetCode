package main

import "strings"

func numUniqueEmails(emails []string) int {
	m := map[string]int{}
	for _, email := range emails {
		se := strings.Split(email, "@")
		if len(se) < 2 {
			continue
		}
		name, domain := se[0], se[1]
		sn := strings.Split(name, "+")
		if len(sn) < 1 {
			continue
		}
		local := strings.ReplaceAll(sn[0], ".", "")
		m[local+"@"+domain]++
	}
	return len(m)
}
