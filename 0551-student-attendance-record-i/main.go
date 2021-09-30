package main

func checkRecord(s string) bool {
	var absent, late int
	for _, r := range s {
		switch r {
		case 'A':
			absent++
			late = 0
			if absent >= 2 {
				return false
			}
		case 'L':
			late++
			if late >= 3 {
				return false
			}
		default:
			late = 0
		}
	}
	return true
}
