package main

import (
	"flag"
	"fmt"
)

var skill = flag.String("skill", " ", "skill help")

func main() {
	// a := func (a int) int{
	// 	return a
	// }

	// fmt.Println(a(10))

	// func (a int){
	// 	fmt.Println(a)
	// 	return
	// }(11)

	//试试把匿名函数传给普通函数
	// visit([] int{1,2,3,4},func(v int){
	// 	fmt.Println(v)
	// })

	//使用匿名函数实现操作封装

	flag.Parse()
	var skil = map[string]func(){
		"w": func() {
			fmt.Println("w")
		},
		"h": func() {
			fmt.Println("h")
		},
	}

	if f, ok := skil[*skill]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
