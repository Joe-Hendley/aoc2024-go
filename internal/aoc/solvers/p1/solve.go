package p1

import (
	"slices"
	"strings"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/integer"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/logger"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/must"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func arrangeInput(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")

	lhs := []int{}
	rhs := []int{}

	for _, line := range lines {
		parts := strings.Fields(line)
		lhnum := must.Atoi(parts[0])
		rhnum := must.Atoi(parts[1])

		lhs = append(lhs, lhnum)
		rhs = append(rhs, rhnum)
	}

	return lhs, rhs
}

func (s *Solver) Part1(input string) int {
	lhs, rhs := arrangeInput(input)

	slices.Sort(lhs)
	slices.Sort(rhs)

	total := 0
	for i := range lhs {
		total += integer.Distance(lhs[i], rhs[i])
	}

	return total
}

func (s *Solver) Part2(input string) int {
	lhs, rhs := arrangeInput(input)

	rhsCounts := make(map[int]int)

	for i := range rhs {
		_, ok := rhsCounts[rhs[i]]
		if ok {
			rhsCounts[rhs[i]] += 1
		} else {
			rhsCounts[rhs[i]] = 1
		}
	}

	total := 0
	for i := range lhs {
		count, ok := rhsCounts[lhs[i]]
		if ok {
			total += (lhs[i] * count)
		}
	}

	return total
}
