package components

import (
	"log"
)

type InitObject interface {
	GameObject
	Init()
}

var _ InitObject = (*InitLogger)(nil)

type InitLogger struct{}

func (s *InitLogger) Init() {
	log.Println("Init called")
}
