package player

import (
	"container/list"
	"github.com/chuckpreslar/emission"
	tm "github.com/nsf/termbox-go"
	"github.com/scottdriscoll/GoInvaders/event"
	"github.com/scottdriscoll/GoInvaders/projectile"
)

const PLAYER_PROJECTILE_SPEED = 30

type Player struct {
	emitter        *emission.Emitter
	minX           int
	maxX           int
	x              int
	y              int
	lastUpdate     int64
	projectileList list.List
}

func (p *Player) Init(emitter *emission.Emitter, minX int, maxX int, y int) {
	p.emitter = emitter
	p.minX = minX
	p.maxX = maxX
	p.y = y
	p.x = int(maxX / 2)

	p.emitter.On(event.EVENT_KEY_LEFT, p.MoveLeft)
	p.emitter.On(event.EVENT_KEY_RIGHT, p.MoveRight)
	p.emitter.On(event.EVENT_KEY_SPACE, p.Fire)
	p.emitter.On(event.EVENT_HEARTBEAT, p.Heartbeat)
	p.emitter.On(event.EVENT_ALIEN_COLLISION, p.DestroyProjectile)
	p.emitter.On(event.EVENT_ALIEN_PROJECTILE_MOVED, p.CheckForCollisions)

	p.DrawPlayer()
}

func (p *Player) DrawPlayer() {
	tm.SetCell(p.x, p.y, '^', tm.ColorDefault, tm.ColorDefault)
}

func (p *Player) ErasePlayer() {
	tm.SetCell(p.x, p.y, ' ', tm.ColorDefault, tm.ColorDefault)
}

func (p *Player) MoveLeft() {
	if p.x-1 > p.minX {
		p.ErasePlayer()
		p.x--
		p.DrawPlayer()
	}
}

func (p *Player) MoveRight() {
	if p.x+1 < p.maxX {
		p.ErasePlayer()
		p.x++
		p.DrawPlayer()
	}
}

func (p *Player) Fire() {
	// Search for a dead projectile to reuse
	for e := p.projectileList.Front(); e != nil; e = e.Next() {
		tempProjectile := e.Value.(*projectile.Projectile)
		if tempProjectile.Alive == false {
			tempProjectile.Init(p.x, p.y-1, p.lastUpdate, PLAYER_PROJECTILE_SPEED, tm.ColorCyan)
			tempProjectile.Draw()

			return
		}

	}
	projectile := new(projectile.Projectile)
	projectile.Init(p.x, p.y-1, p.lastUpdate, PLAYER_PROJECTILE_SPEED, tm.ColorCyan)
	projectile.Draw()
	p.projectileList.PushBack(projectile)
}

func (p *Player) Heartbeat(currentTime int64) {
	p.lastUpdate = currentTime

	for e := p.projectileList.Front(); e != nil; e = e.Next() {
		tempProjectile := e.Value.(*projectile.Projectile)
		if tempProjectile.Alive == true && currentTime > tempProjectile.LastUpdate+tempProjectile.Speed {
			tempProjectile.LastUpdate = currentTime
			tempProjectile.Erase()
			tempProjectile.Y--

			if tempProjectile.Y <= 1 {
				tempProjectile.Alive = false
			} else {
				tempProjectile.Draw()
				p.emitter.Emit(event.EVENT_PLAYER_PROJECTILE_MOVED, tempProjectile.X, tempProjectile.Y)
			}
		}
	}
}

func (p *Player) CheckForCollisions(x int, y int) {
	if p.x == x && p.y == y {
		p.emitter.Emit(event.EVENT_PLAYER_LOSES)
	}
}

func (p *Player) DestroyProjectile(x int, y int) {
	for e := p.projectileList.Front(); e != nil; e = e.Next() {
		tempProjectile := e.Value.(*projectile.Projectile)
		if tempProjectile.Alive == true && tempProjectile.X == x && tempProjectile.Y == y {
			tempProjectile.Erase()
			tempProjectile.Alive = false
		}
	}
}
