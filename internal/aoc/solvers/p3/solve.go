package p3

import (
	"regexp"
	"strings"

	"github.com/Joe-Hendley/aoc2024/internal/aoc/logger"
	"github.com/Joe-Hendley/aoc2024/internal/aoc/must"
)

var mulRegex = regexp.MustCompile(`mul\(\d{0,3}\,\d{0,3}\)`)
var digitRegex = regexp.MustCompile(`\d+`)
var commandRegex = regexp.MustCompile(`(mul\(\d{0,3}\,\d{0,3}\))|(don't\(\))|(do\(\))`)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	muls := mulRegex.FindAllString(input, -1)
	sum := 0
	for _, mul := range muls {
		digits := digitRegex.FindAllString(mul, -1)
		lhs := must.Atoi(digits[0])
		rhs := must.Atoi(digits[1])
		sum += (lhs * rhs)
	}
	return sum
}

func (s *Solver) Part2(input string) int {
	commands := commandRegex.FindAllString(input, -1)
	sum := 0
	doMul := true
	for _, com := range commands {
		switch {
		case strings.HasPrefix(com, "mul") && doMul:
			digits := digitRegex.FindAllString(com, -1)
			lhs := must.Atoi(digits[0])
			rhs := must.Atoi(digits[1])
			sum += (lhs * rhs)
		case com == "do()":
			doMul = true
		case com == "don't()":
			doMul = false
		}
	}
	return sum
}
