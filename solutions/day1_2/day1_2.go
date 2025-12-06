package main

import (
	"advent2025/utils"
	"fmt"
	"strconv"
)

func main() {
	fileLines := utils.ReadDayInputMultiline(1)
	start := 50
	countZeros := 0
	for _, line := range fileLines {
		direction := line[0:1]
		value, _ := strconv.Atoi(line[1:])

		for i := 0; i < value; i++ {
			if direction == "L" {
				start--
			} else {
				start++
			}
			if start < 0 {
				start = start + 100
			}
			start = start % 100
			if start == 0 {
				countZeros++
			}
		}
	}

	fmt.Println(countZeros)
}
