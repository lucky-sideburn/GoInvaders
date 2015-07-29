// GoInvaders project main.go
package main

import (
	"fmt"
	"github.com/chuckpreslar/emission"
	tm "github.com/nsf/termbox-go"
	"github.com/scottdriscoll/GoInvaders/alien"
	"github.com/scottdriscoll/GoInvaders/event"
	"github.com/scottdriscoll/GoInvaders/gameboard"
	"github.com/scottdriscoll/GoInvaders/keyboard"
	"github.com/scottdriscoll/GoInvaders/player"
	"syscall"
	"time"
)

func main() {
	tm.Init()
	tm.Clear(tm.ColorDefault, tm.ColorDefault)

	gameOver := false
	playerWon := false

	emitter := emission.NewEmitter()
	emitter.On(event.EVENT_KEY_CTRL_C, func() {
		gameOver = true
	})

	emitter.On(event.EVENT_PLAYER_WINS, func() {
		playerWon = true
		gameOver = true
	})

	emitter.On(event.EVENT_PLAYER_LOSES, func() {
		gameOver = true
	})

	keyboard := new(keyboard.Keyboard)
	keyboard.Init(emitter)

	mainboard := new(gameboard.Mainboard)
	mainboard.Init(1, 1, 50, 20)

	player := new(player.Player)
	player.Init(emitter, 1, 50, 20)

	alienManager := new(alien.AlienManager)
	alienManager.Init(emitter, 1, 1, 50, 20)

	var currentTime syscall.Timeval

	eventQueue := make(chan tm.Event)
	go func() {
		for {
			eventQueue <- tm.PollEvent()
		}
	}()

	timer := time.NewTimer(time.Duration(time.Millisecond * 60))

	for gameOver == false {
		select {
		case ev := <-eventQueue:
			if ev.Type == tm.EventKey {
				keyboard.HandleEvent(ev.Key)
			}
		case <-timer.C:
			syscall.Gettimeofday(&currentTime)

			emitter.Emit(event.EVENT_HEARTBEAT, (int64(currentTime.Sec)*1e3 + int64(currentTime.Usec)/1e3))
			timer.Reset(time.Duration(time.Millisecond * 60))
		}

		tm.Flush()
	}

	tm.Close()

	if playerWon == true {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose :(")
	}
}
