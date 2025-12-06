package utils

import (
	"fmt"
	"os"
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
