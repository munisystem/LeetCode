package main

func destCity(paths [][]string) string {
	cities := make(map[string]string, len(paths))
	for i := 0; i < len(paths); i++ {
		cities[paths[i][0]] = paths[i][1]
	}
	city := paths[0][0]
	for {
		c, ok := cities[city]
		if !ok {
			break
		}
		city = c
	}
	return city
}
