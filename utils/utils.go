package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadDayInputMultiline(dayNumber int) []string {
	file, err := os.ReadFile(fmt.Sprintf("inputs/day_%d.txt", dayNumber))
	if err != nil {
		panic(err)
	}

	fileLines := strings.Split(string(file), "\n")
	return fileLines
}

func ReadDayInput(dayNumber int) string {
	file, err := os.ReadFile(fmt.Sprintf("inputs/day_%d.txt", dayNumber))
	if err != nil {
		panic(err)
	}

	return string(file)
}

func ToMatrix(input []string) [][]string {
	output := make([][]string, len(input))
	for y, line := range input {
		output[y] = strings.Split(line, "")
	}
	return output
}

func MustToInsts(input []string) []int {
	result := make([]int, len(input))
	for y, line := range input {
		atoi, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		result[y] = atoi
	}
	return result
}

func MustToInt(input string) int {
	atoi, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return atoi
}
