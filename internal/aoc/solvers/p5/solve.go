package p5

import (
	"slices"
	"strings"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/dangerconv"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	rules := make(map[int][]int)
	isParsingRules := true
	sumCorrectMiddles := 0

	for _, line := range strings.Split(input, "\n") {
		switch {
		case line == "":
			isParsingRules = false

		case isParsingRules:
			parts := strings.Split(line, "|")
			x := dangerconv.Atoi(parts[0])
			y := dangerconv.Atoi(parts[1])

			existing, ok := rules[y]
			if ok {
				rules[y] = append(existing, x)
			} else {
				rules[y] = []int{x}
			}

		case !isParsingRules:
			{
				pages := dangerconv.StringSplitToInts(line, ",")
				updateOk := true
				firstIndexes := make(map[int]int, len(pages))

				for index, page := range pages {
					_, ok := firstIndexes[page]
					if !ok {
						firstIndexes[page] = index
					}
				}

				for _, page := range pages {
					requirements, ok := rules[page]
					if ok {
						targetFirstIndex := firstIndexes[page]
						for _, requirement := range requirements {
							requirementFirstIndex, ok := firstIndexes[requirement]
							if ok && requirementFirstIndex > targetFirstIndex {
								updateOk = false
								break
							}
						}
					}
					if !updateOk {
						break
					}
				}

				if updateOk {
					sumCorrectMiddles += pages[(len(pages) / 2)]
				}
			}
		}
	}
	return sumCorrectMiddles
}

func (s *Solver) Part2(input string) int {
	rules := make(map[int][]int)
	isParsingRules := true
	sumCorrectedMiddles := 0

	for _, line := range strings.Split(input, "\n") {
		switch {
		case line == "":
			isParsingRules = false

		case isParsingRules:
			parts := strings.Split(line, "|")
			x := dangerconv.Atoi(parts[0])
			y := dangerconv.Atoi(parts[1])

			existing, ok := rules[y]
			if ok {
				rules[y] = append(existing, x)
			} else {
				rules[y] = []int{x}
			}

		case !isParsingRules:
			{
				pages := dangerconv.StringSplitToInts(line, ",")
				followsRules := true
				firstIndexes := make(map[int]int, len(pages))

				for index, page := range pages {
					_, ok := firstIndexes[page]
					if !ok {
						firstIndexes[page] = index
					}
				}

				for _, page := range pages {
					requirements, ok := rules[page]
					if ok {
						targetFirstIndex := firstIndexes[page]
						for _, requirement := range requirements {
							requirementFirstIndex, ok := firstIndexes[requirement]
							if ok && requirementFirstIndex > targetFirstIndex {
								followsRules = false
								break
							}
						}
					}
					if !followsRules {
						break
					}
				}

				if !followsRules {
					sorted := sortUpdate(rules, pages)
					sumCorrectedMiddles += sorted[(len(sorted) / 2)]
					break
				}
			}
		}
	}
	return sumCorrectedMiddles
}

func sortUpdate(rules map[int][]int, pages []int) []int {
	working := make([]int, len(pages))
	copy(working, pages)
	slices.SortFunc(working, func(a, b int) int {
		aRules, aOk := rules[a]
		bRules, bOk := rules[b]

		switch {
		case !aOk && !bOk:
			return 0

		case bOk && slices.Contains(bRules, a):
			return -1

		case aOk && slices.Contains(aRules, b):
			return 1
		}

		return 0
	})

	return working
}
