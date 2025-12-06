package main

import (
	"advent2025/utils"
)

func main() {
	input := utils.ReadDayInputMultiline(3)

	sum := 0
	for _, line := range input {
		println(line)
		ma := getMaxFromLine(line)
		println(ma)
		sum += ma
	}
	println(sum)
}

func getMaxFromLine(str string) int {
	max1 := 0
	max2 := 0
	firstPos := 0

	for i := 0; i < len(str)-1; i++ {
		if int(str[i]-'0') > max1 {
			max1 = int(str[i] - '0')
			firstPos = i
		}
	}
	for i := firstPos + 1; i < len(str); i++ {
		if int(str[i]-'0') > max2 {
			max2 = int(str[i] - '0')
		}
	}
	return 10*max1 + max2

}
