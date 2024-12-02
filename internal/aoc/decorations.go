package aoc

import (
	"strings"
	"time"
	"unicode/utf8"
)

type decorationPair []string

var decorations = []decorationPair{
	{"📐 ", " 🎯"},
	{"🟥🟩🟦", ""},
	{"🔧", "⚙️"},
	{"🎫", "🎫🎫🎫"},
	{"🌱", "(╯°□°)╯︵ ┻━┻"},
	{"🛥️"},
	{"🐪🐪🐪", "🃏🤡🃏"},
	{"👻", "👻👻👻👻👻👻👻👻"},
	{"🌴"},
	{"🕳️ 🐁"},
	{" 🌌 🌌", "🌌         🌌"},
}

// could make this simpler/neater, but also nah
func decorate(puzzle, part int) string {
	if len(decorations) < puzzle {
		return strings.Repeat(" ", 8)
	}

	if len(decorations[puzzle-1]) < part {
		return strings.Repeat(" ", 8)
	}

	return centre(decorations[puzzle-1][part-1])

}

// fun with runes wasn't what I was planning tonight
func centre(s string) string {
	padded := s
	length := utf8.RuneCountInString(padded)
	if length%2 != 0 {
		padded = padded + " "
		length += 1
	}
	padLength := (8 - length) / 2

	if padLength < 0 {
		padLength = 0
	}

	return strings.Repeat(" ", padLength) + padded + strings.Repeat(" ", padLength)
}

func buildStars(numStars int) string {
	line1 := "\t"
	line2 := "\t"

	for i := 0; i < numStars; i++ {
		switch i % 2 {
		case 0:
			line1 += "🌟"
		case 1:
			line2 += "🌟"
		}
	}

	return line1 + "\n" + line2
}

func buildTimeString(d time.Duration) string {
	if d.Microseconds() < 10 {
		return d.Truncate(time.Nanosecond).String()
	}

	return d.Truncate(time.Microsecond).String()
}
