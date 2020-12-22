package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs[0]) == 0 {
		return ""
	}

	str := strs[0]
	head, tail := 1, len(str)
	for head <= tail {
		mid := (head + tail) / 2
		if isPrefixMatch(strs, str[0:mid]) {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return str[0 : (head+tail)/2]
}

func isPrefixMatch(strs []string, subst string) bool {
	for _, str := range strs {
		if len(subst) > len(str) {
			return false
		}
		if subst != str[0:len(subst)] {
			return false
		}
	}
	return true
}
