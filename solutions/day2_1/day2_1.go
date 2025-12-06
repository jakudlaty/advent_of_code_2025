package main

import (
	"advent2025/utils"
	"strconv"
	"strings"
)

func isDoubledString(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

func main() {

	input := utils.ReadDayInput(2)
	ranges := strings.Split(input, ",")

	sum := 0
	for _, rng := range ranges {
		t := strings.Split(rng, "-")
		start, _ := strconv.Atoi(t[0])
		end, _ := strconv.Atoi(t[1])

		for i := start; i <= end; i++ {
			toTest := strconv.Itoa(i)
			if isDoubledString(toTest) {
				println(toTest)
				sum += i
			}
		}
	}
	println(sum)

}
