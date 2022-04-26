package internal

import "time"

type ticker struct {
	acc      time.Duration
	tickTime time.Duration
}

func (t *ticker) CallOnTick(delta time.Duration, callback func()) {
	t.acc += delta
	if t.acc >= t.tickTime {
		t.acc = 0
		callback()
	}
}
