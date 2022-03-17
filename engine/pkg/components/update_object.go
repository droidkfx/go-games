package components

import (
	"time"
)

type UpdatableObject interface {
	GameObject
	Update(delta time.Duration)
}
