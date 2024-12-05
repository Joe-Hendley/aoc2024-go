package p5

import (
	"slices"
	"strings"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/fun"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/must"
)

func solvePart1Functional(input string) int {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	updates := fun.Map(strings.Split(parts[1], "\n"), func(update string) []int {
		return must.StringSplitToInts(update, ",")
	})
	validUpdates := fun.Filter(updates, func(pages []int) bool {
		return isUpdateValid(rules, pages)
	})

	return fun.Fold(fun.Map(validUpdates, func(update []int) int { return update[len(update)/2] }), func(a, b int) int { return a + b }, 0)
}

func solvePart2Functional(input string) int {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	updates := fun.Map(strings.Split(parts[1], "\n"), func(update string) []int {
		return must.StringSplitToInts(update, ",")
	})
	invalidUpdates := fun.Filter(updates, func(pages []int) bool {
		return !isUpdateValid(rules, pages)
	})

	sortedUpdates := fun.Map(invalidUpdates, func(update []int) []int {
		return sortUpdate(rules, update)
	})

	return fun.Fold(fun.Map(sortedUpdates, func(update []int) int { return update[len(update)/2] }), func(a, b int) int { return a + b }, 0)
}

func parseRules(rules string) map[int][]int {
	parsed := map[int][]int{}
	fun.MapInPlace(strings.Split(rules, "\n"), func(rule string) string {
		split := strings.Split(rule, "|")
		parsed[must.Atoi(split[1])] = append(parsed[must.Atoi(split[1])], must.Atoi(split[0]))
		return rule
	})

	return parsed
}

func isUpdateValid(rules map[int][]int, pages []int) bool {
	firstIndexes := map[int]int{}
	for index, page := range pages {
		if _, ok := firstIndexes[page]; !ok {
			firstIndexes[page] = index
		}
	}

	for _, page := range pages {
		if pageRules, ok := rules[page]; ok {
			for _, rule := range pageRules {
				if ruleFirstIndex, ok := firstIndexes[rule]; ok && ruleFirstIndex > firstIndexes[page] {
					return false
				}
			}
		}
	}

	return true
}

func sortUpdate(rules map[int][]int, update []int) []int {
	working := make([]int, len(update))
	copy(working, update)
	slices.SortFunc(working, func(a, b int) int {
		if aRules, aOk := rules[a]; aOk && slices.Contains(aRules, b) {
			return 1
		}

		if bRules, bOk := rules[b]; bOk && slices.Contains(bRules, a) {
			return -1
		}

		return 0
	})

	return working
}
