package p2

import (
	"slices"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/dangerconv"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/integer"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func isSafe(readings []int) bool {
	safe := true
	previousDirection := readings[1] < readings[0]
	for i := 1; i < len(readings); i++ {
		distance := integer.Distance(readings[i], readings[i-1])
		direction := readings[i] < readings[i-1]
		if distance < 1 || distance > 3 || direction != previousDirection {
			safe = false
			break
		}
		previousDirection = direction
	}

	return safe
}

func (s *Solver) Part1(input string) int {
	reports := dangerconv.StringToLinesOfInts(input)
	safeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func (s *Solver) Part2(input string) int {
	reports := dangerconv.StringToLinesOfInts(input)
	safeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		} else {
			for i := range report {
				if isSafe(slices.Concat(report[:i], report[i+1:])) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}
