package p4

import (
	"github.com/Joe-Hendley/aoc2024/internal/aoc/grid"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/grid/direction"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	xmasGrid := grid.FromString(input)

	total := 0
	searchStr := []rune("XMAS")

	for y := range xmasGrid.Height() {
		for x := range xmasGrid.Width() {
			if xmasGrid.At(x, y) == searchStr[0] {
				for _, direction := range direction.All() {
					if xmasGrid.CheckCellsInDirection(searchStr, direction, x, y) {
						total++
					}
				}
			}
		}
	}

	return total
}

func (s *Solver) Part2(input string) int {
	xmasGrid := grid.FromString(input)

	total := 0

	for y := 1; y < xmasGrid.Height()-1; y++ {
		for x := 1; x < xmasGrid.Width()-1; x++ {
			if xmasGrid.At(x, y) == 'A' {
				count := 0
				for _, direction := range direction.Diagonal() {
					if xmasGrid.CheckCellInDirection('M', direction, x, y) && xmasGrid.CheckCellInDirection('S', direction.Opposite(), x, y) {
						count++
					}
				}
				if count == 2 {
					total++
				}
			}
		}
	}

	return total
}
