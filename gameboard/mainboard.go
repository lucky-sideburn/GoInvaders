package gameboard

import (
	tm "github.com/nsf/termbox-go"
)

type Mainboard struct {
	xStart int
	yStart int
	width  int
	height int
}

func (m *Mainboard) Init(xStart int, yStart int, width int, height int) {
	m.xStart = xStart
	m.yStart = yStart
	m.width = width
	m.height = height

	var x int = xStart + 1
	var y int = yStart + 1

	// Draw Top
	tm.SetCell(xStart, yStart, '+', tm.ColorDefault, tm.ColorDefault)
	for x < xStart+width {
		tm.SetCell(x, yStart, '-', tm.ColorDefault, tm.ColorDefault)
		x++
	}
	tm.SetCell(x, yStart, '+', tm.ColorDefault, tm.ColorDefault)

	// Draw Sides
	for y < yStart+height {
		tm.SetCell(xStart, y, '|', tm.ColorDefault, tm.ColorDefault)
		tm.SetCell(xStart+width, y, '|', tm.ColorDefault, tm.ColorDefault)
		y++
	}

	// Draw Bottom
	x = xStart + 1
	tm.SetCell(xStart, y, '+', tm.ColorDefault, tm.ColorDefault)
	for x < xStart+width {
		tm.SetCell(x, y, '-', tm.ColorDefault, tm.ColorDefault)
		x++
	}
	tm.SetCell(x, y, '+', tm.ColorDefault, tm.ColorDefault)
}
