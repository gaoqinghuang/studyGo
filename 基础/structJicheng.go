package main

import (
	"fmt"
)

type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("I can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println("I can walk")
}

type Human struct {
	Walkable
}
type Bird struct {
	Walkable
	Flying
}

func main() {
	b := new(Bird)
	fmt.Println("Bird:")
	b.Fly()
	b.Walk()
	c := new(Human)
	fmt.Println("Human:")
	c.Walk()
}
