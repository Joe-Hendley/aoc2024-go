package p4

import (
	"strings"
)

type xmasSearch struct {
	letters string
	found   []bool
	width   int
	height  int
}

func newXmasSearch(input string) xmasSearch {
	width := strings.Index(input, "\n")
	letters := strings.ReplaceAll(input, "\n", "")
	height := len(letters) / width

	return xmasSearch{
		letters: letters,
		found:   make([]bool, len(letters)),
		width:   width,
		height:  height,
	}
}

func (ws xmasSearch) at(x, y int) string {
	if x < 0 || y < 0 || x >= ws.width || y >= ws.height {
		return ""
	}

	return string(ws.letters[ws.idx(x, y)])
}

func (ws xmasSearch) idx(x, y int) int {
	return (y * ws.width) + x
}

func (ws xmasSearch) checkUpForXMAS(x, y int) bool {
	if y < 1 || y > ws.height-2 || x < 1 || x > ws.width-2 {
		return false
	}

	if ws.at(x, y) != "A" {
		return false
	}

	if ws.at(x-1, y-1) == "M" && ws.at(x+1, y-1) == "M" && ws.at(x-1, y+1) == "S" && ws.at(x+1, y+1) == "S" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws xmasSearch) checkDownForXMAS(x, y int) bool {
	if y < 1 || y > ws.height-2 || x < 1 || x > ws.width-2 {
		return false
	}

	if ws.at(x, y) != "A" {
		return false
	}

	if ws.at(x-1, y-1) == "S" && ws.at(x+1, y-1) == "S" && ws.at(x-1, y+1) == "M" && ws.at(x+1, y+1) == "M" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws xmasSearch) checkLeftForXMAS(x, y int) bool {
	if y < 1 || y > ws.height-2 || x < 1 || x > ws.width-2 {
		return false
	}

	if ws.at(x, y) != "A" {
		return false
	}

	if ws.at(x-1, y-1) == "M" && ws.at(x+1, y-1) == "S" && ws.at(x-1, y+1) == "M" && ws.at(x+1, y+1) == "S" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws xmasSearch) checkRightForXMAS(x, y int) bool {
	if y < 1 || y > ws.height-2 || x < 1 || x > ws.width-2 {
		return false
	}

	if ws.at(x, y) != "A" {
		return false
	}

	if ws.at(x-1, y-1) == "S" && ws.at(x+1, y-1) == "M" && ws.at(x-1, y+1) == "S" && ws.at(x+1, y+1) == "M" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws xmasSearch) String() string {
	s := ""
	for y := range ws.height {
		for x := range ws.width {
			if ws.found[ws.idx(x, y)] {
				s += ws.at(x, y)
			} else {
				s += "."
			}
		}
		s += "\n"
	}

	return s
}
