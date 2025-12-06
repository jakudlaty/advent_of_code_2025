package main

import (
	"advent2025/utils"
	"strconv"
	"strings"
)

func isRepeatedStrings(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		pat := s[0:i]
		matc := strings.Repeat(pat, len(s)/len(pat))
		if matc == s {
			println(s)
			return true
		}
	}
	return false
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
			if isRepeatedStrings(toTest) {
				sum += i
			}
		}
	}
	println(sum)
}
