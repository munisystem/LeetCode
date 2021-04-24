package main

func reverseOnlyLetters(S string) string {
	i, j := 0, len(S)-1
	ans := make([]byte, len(S), len(S))
	for i <= j {
		if !((S[i] >= 65 && S[i] <= 90) || (S[i] >= 97 && S[i] <= 122)) {
			ans[i] = S[i]
			i++
			continue
		}
		if !((S[j] >= 65 && S[j] <= 90) || (S[j] >= 97 && S[j] <= 122)) {
			ans[j] = S[j]
			j--
			continue
		}
		ans[i], ans[j] = S[j], S[i]
		i++
		j--
	}
	return string(ans)
}
