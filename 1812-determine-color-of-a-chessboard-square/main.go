package main

func squareIsWhite(coordinates string) bool {
	col, row := int(coordinates[0])-'a', int(coordinates[1])-'1'
	return col%2 != row%2
}
