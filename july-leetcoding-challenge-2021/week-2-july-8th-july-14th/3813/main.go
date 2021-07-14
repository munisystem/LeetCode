package main

func customSortString(order string, str string) string {
	var ans []byte
	chars := map[int]int{}
	for i := 0; i < len(str); i++ {
		chars[int(str[i])-'a']++
	}
	for i := 0; i < len(order); i++ {
		for chars[int(order[i])-'a'] > 0 {
			ans = append(ans, order[i])
			chars[int(order[i])-'a']--
		}
	}
	for k, v := range chars {
		for i := 0; i < v; i++ {
			ans = append(ans, byte(k+'a'))
		}
	}
	return string(ans)
}
