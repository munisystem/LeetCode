package main

import "strings"

func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	arr := make([]string, len(strs), len(strs))
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		arr[str[len(str)-1]-'1'] = str[:len(str)-1]
	}
	return strings.Join(arr, " ")
}
