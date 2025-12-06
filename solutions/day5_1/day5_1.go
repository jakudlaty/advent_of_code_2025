package main

import (
	"advent2025/utils"
	"fmt"
	"strings"
)

func main() {
	ranges, numbersToTest := ReadInputForDay5()
	fresh := 0
	for _, num := range numbersToTest {
		if InRanges(ranges, num) {
			fresh++
		}
	}
	fmt.Println(fresh)
}

func ReadInputForDay5() ([][]int, []int) {
	var lines = utils.ReadDayInputMultiline(5)
	var ranges [][]int
	var numbersToTest []int

	for _, line := range lines {
		if line == "" {
			continue
		}
		res := strings.Split(line, "-")
		if len(res) == 2 {
			insts := utils.MustToInsts(res)
			ranges = append(ranges, insts[0:2])
		}
		if len(res) == 1 {
			num := utils.MustToInt(res[0])
			numbersToTest = append(numbersToTest, num)
		}
	}
	return ranges, numbersToTest
}

func InRanges(ranges [][]int, value int) bool {
	for _, ra := range ranges {
		if value >= ra[0] && value <= ra[1] {
			return true
		}
	}
	return false
}
