package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data,_ := ioutil.ReadFile("text.txt")
	//err := ioutil.WriteFile("text.txt",[]byte("我"),666)
	fmt.Println(string(data))
}
