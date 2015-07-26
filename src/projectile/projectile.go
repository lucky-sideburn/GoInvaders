package projectile

import (
    tm "github.com/nsf/termbox-go"
)

type Projectile struct {
	X int
    Y int
    LastUpdate int64
    Speed int64
    Alive bool
    color tm.Attribute
}

func (p *Projectile) Init(x int, y int, lastUpdate int64, speed int64, color tm.Attribute) {
    p.X = x
    p.Y = y
    p.LastUpdate = lastUpdate
    p.Speed = speed
    p.Alive = true
    p.color = color
}

func (p *Projectile) Erase() {
    tm.SetCell(p.X, p.Y, ' ', tm.ColorDefault, tm.ColorDefault)
}

func (p *Projectile) Draw() {
    tm.SetCell(p.X, p.Y, '|', p.color, tm.ColorDefault)
}
