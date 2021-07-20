package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var un [][]uint8
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			un[i][j] = uint8((dx+j)*i + j/2)
		}
	}
	return un

}

func main() {
	pic.Show(Pic(10, 11))
}
