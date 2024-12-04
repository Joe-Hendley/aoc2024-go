package p4

import (
	"testing"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 18
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 9
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}