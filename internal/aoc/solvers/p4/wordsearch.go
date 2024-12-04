package p4

import "strings"

type wordSearch struct {
	letters string
	found   []bool
	width   int
	height  int
}

func newWordSearch(input string) wordSearch {
	width := strings.Index(input, "\n")
	letters := strings.ReplaceAll(input, "\n", "")
	height := len(letters) / width

	return wordSearch{
		letters: letters,
		found:   make([]bool, len(letters)),
		width:   width,
		height:  height,
	}
}

func (ws wordSearch) at(x, y int) string {
	if x < 0 || y < 0 || x >= ws.width || y >= ws.height {
		return ""
	}

	return string(ws.letters[ws.idx(x, y)])
}

func (ws wordSearch) idx(x, y int) int {
	return (y * ws.width) + x
}

func (ws wordSearch) checkUpForXMAS(x, y int) bool {
	if y < 2 {
		return false
	}

	if ws.at(x, y)+ws.at(x, y-1)+ws.at(x, y-2)+ws.at(x, y-3) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkDownForXMAS(x, y int) bool {
	if y > ws.height-4 {
		return false
	}

	if ws.at(x, y)+ws.at(x, y+1)+ws.at(x, y+2)+ws.at(x, y+3) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkLeftForXMAS(x, y int) bool {
	if x < 2 {
		return false
	}

	if ws.at(x, y)+ws.at(x-1, y)+ws.at(x-2, y)+ws.at(x-3, y) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkRightForXMAS(x, y int) bool {
	if x > ws.width-4 {
		return false
	}

	if ws.at(x, y)+ws.at(x+1, y)+ws.at(x+2, y)+ws.at(x+3, y) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkUpLeftForXMAS(x, y int) bool {
	if y < 2 || x < 2 {
		return false
	}

	if ws.at(x, y)+ws.at(x-1, y-1)+ws.at(x-2, y-2)+ws.at(x-3, y-3) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkUpRightForXMAS(x, y int) bool {
	if y < 2 || x > ws.width-4 {
		return false
	}

	if ws.at(x, y)+ws.at(x+1, y-1)+ws.at(x+2, y-2)+ws.at(x+3, y-3) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkDownLeftForXMAS(x, y int) bool {
	if y > ws.height-4 || x < 2 {
		return false
	}

	if ws.at(x, y)+ws.at(x-1, y+1)+ws.at(x-2, y+2)+ws.at(x-3, y+3) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) checkDownRightForXMAS(x, y int) bool {
	if y > ws.height-4 || x > ws.width-4 {
		return false
	}

	if ws.at(x, y)+ws.at(x+1, y+1)+ws.at(x+2, y+2)+ws.at(x+3, y+3) == "XMAS" {
		ws.found[ws.idx(x, y)] = true
		return true
	}

	return false
}

func (ws wordSearch) String() string {
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
