package p3

import (
	"testing"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.DangerReadToString("test.txt")
	want := 161
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

const testPartTwo = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestPartTwo(t *testing.T) {
	//input := file.DangerReadToString("test.txt")
	input := testPartTwo
	want := 48
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}
