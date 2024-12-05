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
	printQueue := newPrintQueue(input)
	return printQueue.sumMiddleOfCorrectUpdates()
}

func (s *Solver) Part2(input string) int {
	printQueue := newPrintQueue(input)
	return printQueue.sumMiddleOfCorrectedUpdates()
}

type printQueue struct {
	rules   map[int][]int
	updates [][]int
}

func newPrintQueue(input string) printQueue {
	parts := strings.Split(input, "\n\n")
	rules := make(map[int][]int, len(parts[0]))

	for _, line := range strings.Split(parts[0], "\n") {
		ruleParts := strings.Split(line, "|")
		x := dangerconv.Atoi(ruleParts[0])
		y := dangerconv.Atoi(ruleParts[1])

		existing, ok := rules[y]
		if ok {
			rules[y] = append(existing, x)
		} else {
			rules[y] = []int{x}
		}
	}

	updates := make([][]int, 0, len(parts[1]))
	for _, line := range strings.Split(parts[1], "\n") {
		update := dangerconv.StringSplitToInts(line, ",")
		updates = append(updates, update)
	}

	return printQueue{
		rules:   rules,
		updates: updates,
	}
}

func (pq printQueue) sortUpdate(update []int) []int {
	working := make([]int, len(update))
	copy(working, update)
	slices.SortFunc(working, func(a, b int) int {
		aRules, aOk := pq.rules[a]
		bRules, bOk := pq.rules[b]

		switch {
		case bOk && slices.Contains(bRules, a):
			return -1

		case aOk && slices.Contains(aRules, b):
			return 1
		}

		return 0
	})

	return working
}

func (pq printQueue) sumMiddleOfCorrectUpdates() int {
	sumCorrectMiddles := 0

	for _, update := range pq.updates {
		followsRules := true
		firstIndexes := make(map[int]int, len(update))

		for index, page := range update {
			_, ok := firstIndexes[page]
			if !ok {
				firstIndexes[page] = index
			}
		}

		for _, page := range update {
			requirements, ok := pq.rules[page]
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

		if followsRules {
			sumCorrectMiddles += update[(len(update) / 2)]
		}
	}

	return sumCorrectMiddles
}

func (pq printQueue) sumMiddleOfCorrectedUpdates() int {
	sumCorrectedMiddles := 0

	for _, update := range pq.updates {
		followsRules := true
		firstIndexes := make(map[int]int, len(update))

		for index, page := range update {
			_, ok := firstIndexes[page]
			if !ok {
				firstIndexes[page] = index
			}
		}

		for _, page := range update {
			requirements, ok := pq.rules[page]
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
			sorted := pq.sortUpdate(update)
			sumCorrectedMiddles += sorted[(len(sorted) / 2)]
		}
	}

	return sumCorrectedMiddles
}
