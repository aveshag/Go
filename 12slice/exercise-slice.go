package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var grid [][]uint8

	for i := 0; i < dx; i++ {
		row := []uint8{}
		for j := 0; j < dy; j++ {
			row = append(row, uint8(i^j))
		}
		grid = append(grid, row)
	}

	return grid
}

func main() {
	pic.Show(Pic)
}
