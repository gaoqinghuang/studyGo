package main

import (
	"bytes"
	"fmt"
)

func joinString(slist ...string) string {
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint("www", slist)
}

func main() {
	// fmt.Println(joinString("pig"," and"," pig"))
	// s := 1
	//    str := fmt.Sprintf("%v", s)
	print(1, 2, 3)
}
