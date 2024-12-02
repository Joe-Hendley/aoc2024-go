package dangerconv

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func StringFieldsToInts(s string) []int {
	fields := strings.Fields(s)
	ints := make([]int, len(fields))

	for i := range fields {
		ints[i] = Atoi(fields[i])
	}

	return ints
}

func StringSplitToInts(s string, sep string) []int {
	split := strings.Split(s, sep)
	ints := make([]int, len(split))

	for i := range split {
		ints[i] = Atoi(split[i])
	}

	return ints
}
