package main

import (
	"fmt"
)

var a [10]int
var b = [...]int{1, 2, 3}

func main() {
	// fmt.Println(a,b)

	//遍历
	for k, v := range b {
		fmt.Println(k, v)
	}
}
