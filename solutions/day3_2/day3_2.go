package main

import (
	"advent2025/utils"
	"strconv"
)

func main() {
	input := utils.ReadDayInputMultiline(3)

	sum := 0
	for _, line := range input {
		println(line)
		ma, _ := getMaxFromLine(line)
		println(ma)
		sum += ma
	}
	println(sum)
}

func getMaxFromLine(str string) (int, error) {
	ints := getMaxFromLineCapped(str, 12, 0, 0)
	out := ""
	for _, i := range ints {
		out += strconv.Itoa(i)
	}
	return strconv.Atoi(out)
}

func getMaxFromLineCapped(str string, leaveCharsOnRight int, startPos int, charsFound int) []int {
	if charsFound == 12 {
		return []int{}
	}

	foundPos := -1
	foundVal := -1
	for i := 9; i >= 0; i-- {
		for x := startPos; x <= len(str)-leaveCharsOnRight; x++ {
			atoi, err := strconv.Atoi(str[x : x+1])
			if err != nil {
				panic(err)
			}
			if atoi == i {
				foundPos = x
				foundVal = atoi
				break
			}
		}
		if foundPos != -1 {
			break
		}
	}
	rest := getMaxFromLineCapped(str, leaveCharsOnRight-1, foundPos+1, charsFound+1)
	return append([]int{foundVal}, rest...)
}
