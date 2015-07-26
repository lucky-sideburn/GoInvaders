package alien

import (
    "container/list"
    tm "github.com/nsf/termbox-go"
    "github.com/chuckpreslar/emission"
    "event"
    "math/rand"
    "projectile"
)

const DIRECTION_LEFT = 1
const DIRECTION_RIGHT = 2
const ALIEN_SPEED = 150
const ALIEN_PROJECTILE_SPEED = 100
const ALIEN_STARTING_FIRE_CHANCE = 40

type AlienManager struct {
    emitter *emission.Emitter
    currentDirection int
    minX int
    minY int
    maxX int
    maxY int
    aliensLeft int
    alienFireChance int
    currentAlienSpeed int64
    lastUpdate int64
    alienList list.List
    alienProjectileList list.List
}

func (a *AlienManager) Init(emitter *emission.Emitter, minX int, minY int, maxX int, maxY int) {
    a.emitter = emitter
    a.currentDirection = DIRECTION_RIGHT
    a.minX = minX
    a.minY = minY
    a.maxX = maxX
    a.maxY = maxY
    a.lastUpdate = 0
    a.aliensLeft = 0
    a.currentAlienSpeed = ALIEN_SPEED
    a.alienFireChance = ALIEN_STARTING_FIRE_CHANCE
    
    for x := minX + 12; x <= maxX - 12; x++ {
        for y := minY + 1; y <= minY + 6; y++ {
            alien := new(Alien)
            alien.Init(x, y)
            alien.Draw()
            a.alienList.PushBack(alien)
            a.aliensLeft++
        }
    }

    a.emitter.On(event.EVENT_HEARTBEAT, a.Heartbeat)
    a.emitter.On(event.EVENT_PLAYER_PROJECTILE_MOVED, a.TestForCollisions)
}

func (a *AlienManager) Heartbeat(currentTime int64) {
    if a.lastUpdate == 0 {
        a.lastUpdate = currentTime
    }
    
    if currentTime > a.lastUpdate + a.currentAlienSpeed {
        a.MoveAliens()
        a.lastUpdate = currentTime
    }
    
    a.UpdateProjectiles(currentTime)
}

func (a *AlienManager) MoveAliens() {
    changeDirection := false
    
    for e := a.alienList.Front(); e != nil; e = e.Next() {
        tempAlien := e.Value.(*Alien)
        if tempAlien.Alive == false {
            continue
        }
        tempAlien.Erase()
        if a.currentDirection == DIRECTION_LEFT {
            tempAlien.X--;
            if tempAlien.X == a.minX + 1 {
                changeDirection = true
            }
        } else {
            tempAlien.X++;
            if tempAlien.X == a.maxX{
                changeDirection = true
            }
        }
    }

    if changeDirection == true {
        for e := a.alienList.Front(); e != nil; e = e.Next() {
            tempAlien := e.Value.(*Alien)
            if tempAlien.Alive == true {
                tempAlien.Y++;
                if tempAlien.Y == a.maxY {
                    a.emitter.Emit(event.EVENT_PLAYER_LOSES)
                }
            }
        }
    }

    for e := a.alienList.Front(); e != nil; e = e.Next() {
        tempAlien := e.Value.(*Alien)
        if tempAlien.Alive == true {
            tempAlien.Draw()

            if rand.Intn(10000) <= a.alienFireChance {
                a.Fire(tempAlien.X, tempAlien.Y + 1, a.lastUpdate)
            }
        }
    }
    
    if (changeDirection == true) {
        if a.currentDirection == DIRECTION_LEFT {
            a.currentDirection = DIRECTION_RIGHT
        } else {
            a.currentDirection = DIRECTION_LEFT
        }
    }
}

func (a *AlienManager) UpdateProjectiles(currentTime int64) {
    for e := a.alienProjectileList.Front(); e != nil; e = e.Next() {
        tempProjectile := e.Value.(*projectile.Projectile)
        if tempProjectile.Alive == true && currentTime > tempProjectile.LastUpdate + tempProjectile.Speed {
            tempProjectile.LastUpdate = currentTime
            tempProjectile.Erase()
            tempProjectile.Y++;
            
            if tempProjectile.Y >= a.maxY + 1 {
                tempProjectile.Alive = false
            } else {
                tempProjectile.Draw()
                a.emitter.Emit(event.EVENT_ALIEN_PROJECTILE_MOVED, tempProjectile.X, tempProjectile.Y)
            }
        }
    }
}

func (a *AlienManager) Fire(x int, y int, lastUpdate int64) {
    // Search for a dead projectile to reuse
    for e := a.alienProjectileList.Front(); e != nil; e = e.Next() {
        tempProjectile := e.Value.(*projectile.Projectile)
        if (tempProjectile.Alive == false) {
            tempProjectile.Init(x, y, a.lastUpdate, ALIEN_PROJECTILE_SPEED, tm.ColorRed)
            tempProjectile.Draw()
            
            return
        }
        
    }
    
    projectile := new(projectile.Projectile)
    projectile.Init(x, y, a.lastUpdate, ALIEN_PROJECTILE_SPEED, tm.ColorRed)   
    projectile.Draw()
    a.alienProjectileList.PushBack(projectile)
}

func (a *AlienManager) TestForCollisions(x int, y int) {
    for e := a.alienList.Front(); e != nil; e = e.Next() {
        tempAlien := e.Value.(*Alien)
        
        if tempAlien.Alive == true && tempAlien.X == x && tempAlien.Y == y {
            tempAlien.Erase()
            tempAlien.Alive = false
            a.aliensLeft--
            a.currentAlienSpeed--
            a.alienFireChance += 2
            a.emitter.Emit(event.EVENT_ALIEN_COLLISION, x, y)
            if a.aliensLeft == 0 {
                a.emitter.Emit(event.EVENT_PLAYER_WINS)
            }
            
            return
        }
    }
}
