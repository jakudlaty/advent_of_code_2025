package main

import (
	"advent2025/utils"
)

func IsOccupied(x1 int, y1 int, pile [][]string) bool {
	if x1 < 0 || y1 < 0 {
		return false
	}
	if x1 >= len(pile[0]) {
		return false
	}
	if y1 >= len(pile) {
		return false
	}
	u := pile[y1][x1]
	return u == "@"
}

func CanBeHandled(lines [][]string, x int, y int) bool {
	occupied := 0
	for x1 := x - 1; x1 <= x+1; x1++ {
		for y1 := y - 1; y1 <= y+1; y1++ {
			//skip our place
			if x == x1 && y == y1 {
				continue
			}

			if IsOccupied(x1, y1, lines) {
				occupied++
			}
		}
	}
	return occupied < 4 && lines[y][x] == "@"
}

func main() {
	var lines = utils.ReadDayInputMultiline(4)
	matrix := utils.ToMatrix(lines)
	count := 0
	for y, line := range matrix {
		for x, _ := range line {
			if CanBeHandled(matrix, x, y) {
				count++
			}
		}
	}
	println(count)
}
