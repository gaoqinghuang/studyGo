package main

import "fmt"

func main() {

	//从标准输入中获取
	a := ""
	fmt.Println(fmt.Scanln(&a))
	//格式化
	fmt.Println(fmt.Sprintf("我%d", 12))
	//打印
	fmt.Println(11)
}
