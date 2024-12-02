package p1

import (
	"testing"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 11
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 31
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}