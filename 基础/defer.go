package main

import (
	"fmt"
	"os"
)

//延迟执行语句
func main() {

	size := fileSize("99.go")
	fmt.Println(size)
}

//适用场景,打开文件，自动关闭文件
func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}
