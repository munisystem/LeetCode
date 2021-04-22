package main

func flipAndInvertImage(image [][]int) [][]int {
	cols, rows := len(image), len(image[0])
	for i := 0; i < cols; i++ {
		for j := 0; j < (rows+1)/2; j++ {
			image[i][j], image[i][rows-j-1] = image[i][rows-j-1]^1, image[i][j]^1
		}
	}
	return image
}
