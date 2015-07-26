package alien

import (
    tm "github.com/nsf/termbox-go"
)

type Alien struct {
    X int
    Y int
    Alive bool
}

func (a *Alien) Init(x int, y int) {
    a.X = x
    a.Y = y
    a.Alive = true
}

func (a *Alien) Draw() {
    tm.SetCell(a.X, a.Y, 'M', tm.ColorMagenta, tm.ColorDefault)
}

func (a *Alien) Erase() {
    tm.SetCell(a.X, a.Y, ' ', tm.ColorDefault, tm.ColorDefault)
}
