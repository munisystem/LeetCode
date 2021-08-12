package main

func groupAnagrams(strs []string) [][]string {
	m := map[[26]byte][]string{}
	for i := 0; i < len(strs); i++ {
		var chars [26]byte
		for j := 0; j < len(strs[i]); j++ {
			chars[int(strs[i][j])-'a']++
		}
		m[chars] = append(m[chars], strs[i])
	}
	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
