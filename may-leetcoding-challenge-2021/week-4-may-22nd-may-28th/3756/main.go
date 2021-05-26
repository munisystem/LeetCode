package main

func minPartitions(n string) int {
	var ans byte
	for i := 0; i < len(n); i++ {
		num := n[i]
		if ans < num {
			ans = num
		}
	}
	return int(ans - 48)
}
