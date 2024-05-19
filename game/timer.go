package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currentTicks int
	tagetTicks   int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		tagetTicks:   int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.tagetTicks {
		t.currentTicks++
	}
}

func (t *Timer) isReady() bool {
	return t.currentTicks >= t.tagetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}
