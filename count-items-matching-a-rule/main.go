package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var rule int
	switch ruleKey {
	case "type":
		rule = 0
	case "color":
		rule = 1
	case "name":
		rule = 2
	}
	ans := 0
	for i := 0; i < len(items); i++ {
		if items[i][rule] == ruleValue {
			ans++
		}
	}
	return ans
}
