package main

import "math"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	times := make([]int, 26, 26)
	for i := 0; i < len(releaseTimes); i++ {
		if i == 0 {
			times[keysPressed[i]-'a'] = releaseTimes[i]
		} else {
			times[keysPressed[i]-'a'] = int(math.Max(float64(releaseTimes[i]-releaseTimes[i-1]), float64(times[keysPressed[i]-'a'])))
		}
	}
	var ans byte
	max := math.MinInt32
	for i := 0; i < len(times); i++ {
		if max <= times[i] {
			max = times[i]
			ans = byte('a' + i)
		}
	}
	return ans
}
