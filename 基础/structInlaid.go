package main

import (
	"fmt"
)

type BasicColor struct {
	R, G, B float32
}

type Color struct {
	BasicColor
	Alpha float32
	R     float32
}

// func main() {
// 	var c Color
// 	c.BasicColor.R=1
// 	c.R = 2
// 	c.G = 1
// 	c.B = 0

// 	c.Alpha = 1

// 	fmt.Printf("%+v",c)
// }

//匿名结构体的定义

type Wheel struct {
	Size int
}

type Car struct {
	Wheel
	Engine struct {
		Power int
		Type  string
	}
}

func main() {
	car := Car{
		Wheel: Wheel{
			Size: 10,
		},
		Engine: struct {
			Power int
			Type  string
		}{
			Power: 143,
			Type:  "1.4T",
		},
	}
	fmt.Println(car)
	// fmt.Printf("%+v",car)
}
