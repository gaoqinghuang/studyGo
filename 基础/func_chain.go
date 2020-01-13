package main

import (
	"fmt"
	"strings"
)

// func StringProccess(list []string,chain []func(string) string) {
func StringProccess(list []string) {
	for index, str := range list {
		result := str
		// for _,proc := range chain {
		// 	result = proc(result)
		// }
		result = removePrefix(result)
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// chain := []func(string) string{
	// 	removePrefix,
	// 	strings.TrimSpace,
	// 	strings.ToUpper,
	// }

	// StringProccess(list,chain)
	StringProccess(list) //切片传递的是指针，所以会被改变，如果传的是数组则不会

	for _, str := range list {
		fmt.Println(str)
	}
}
