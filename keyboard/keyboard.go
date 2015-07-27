package keyboard

import (
    tm "github.com/nsf/termbox-go"
    "github.com/chuckpreslar/emission"
    "github.com/scottdriscoll/GoInvaders/event"
)

type Keyboard struct {
    emitter *emission.Emitter
    eventQueue chan tm.Event
}

func (k *Keyboard) Init(emitter *emission.Emitter) {
    k.emitter = emitter
}

func (k *Keyboard) HandleEvent(key tm.Key) {
    switch key {
    case tm.KeyCtrlC:
        k.emitter.Emit(event.EVENT_KEY_CTRL_C)
    case tm.KeyArrowLeft:
        k.emitter.Emit(event.EVENT_KEY_LEFT)
    case tm.KeyArrowRight:
        k.emitter.Emit(event.EVENT_KEY_RIGHT)
    case tm.KeySpace:
        k.emitter.Emit(event.EVENT_KEY_SPACE)
    }
}
