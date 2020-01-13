package main

import (
	"fmt"
)

type Service interface {
	Start()
	Log(string)
}

type Logger struct{}

func (g *Logger) Log(l string) {
	fmt.Println("我是log")
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {

}

func main() {
	var s Service = new(GameService)
	// var service Service
	// service = s
	s.Start()
	s.Log("hello")
}
