package main

import (
	"fmt"
)

func main() {
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = false
	fmt.Println(any)

	var a int = 1
	var b interface{} = a
	// var c int = b.(int)
	fmt.Println(a == b)
}
