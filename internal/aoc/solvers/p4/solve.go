package p4

import (
	"github.com/Joe-Hendley/aoc2024/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	wordSearch := newWordSearch(input)

	total := 0

	for y := range wordSearch.height {
		for x := range wordSearch.width {
			if wordSearch.at(x, y) == "X" {
				for _, f := range []func(x, y int) bool{
					wordSearch.checkUpForXMAS,
					wordSearch.checkDownForXMAS,
					wordSearch.checkLeftForXMAS,
					wordSearch.checkRightForXMAS,
					wordSearch.checkUpLeftForXMAS,
					wordSearch.checkUpRightForXMAS,
					wordSearch.checkDownLeftForXMAS,
					wordSearch.checkDownRightForXMAS,
				} {
					if f(x, y) {
						total++
					}
				}
			}
		}
	}

	return total
}

func (s *Solver) Part2(input string) int {
	xmasSearch := newXmasSearch(input)

	total := 0

	for y := 1; y < xmasSearch.height-1; y++ {
		for x := 1; x < xmasSearch.width-1; x++ {
			if xmasSearch.at(x, y) == "A" {
				for _, f := range []func(x, y int) bool{
					xmasSearch.checkUpForXMAS,
					xmasSearch.checkDownForXMAS,
					xmasSearch.checkLeftForXMAS,
					xmasSearch.checkRightForXMAS,
				} {
					if f(x, y) {
						total++
					}
				}
			}
		}
	}

	return total
}
