package main

import (
	"fmt"
	"strconv"
)

func main() {
	//总体没看到什么有用的，就下面3个经常用到

	//数字转字符串
	fmt.Println(strconv.Itoa(10))
	//字符串转数字
	fmt.Println(strconv.Atoi("12"))

	//int或者uint的位数
	//fmt.Println(strconv.IntSize)
}

