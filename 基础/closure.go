package main

import (
	"fmt"
)

func main() {
	// str := "hello world"
	// foo := func(){
	// 	str = "hello dude"
	// }
	// foo()
	// fmt.Println(str)

	accumulate := Accumulate(1)
	fmt.Println(accumulate())
	fmt.Println(accumulate())
	fmt.Printf("%p\n", accumulate)
	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)
	fmt.Println(accumulate())

}

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
