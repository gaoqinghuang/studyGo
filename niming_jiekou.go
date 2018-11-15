package main

import (
	"fmt"
)

type Invoker interface {
	Call(interface{})
}

type Struct struct{

}

func (s *Struct) Call(p interface{}){
	fmt.Println("from struct",p)
}

type FuncCaller func (interface{})

func (f FuncCaller) Call(p interface{}){
	f(p)
}

//看得不是很懂，等看到接口的时候再来回头看，@todo
func main() {
	var invoker Invoker
	s := new(Struct)

	invoker = s
	invoker.Call("hello")

	invoker = FuncCaller(func(v interface{}){
		fmt.Println("from function", v)
	})

	invoker.Call("hello")
}