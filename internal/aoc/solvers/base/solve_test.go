package base

import (
	"testing"

	"github.com/Joe-Hendley/aoc2023/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2023/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 0
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 0
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}
