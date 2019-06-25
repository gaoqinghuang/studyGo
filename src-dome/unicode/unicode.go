package main

import (
	"fmt"
	"unicode"
)

func main() {
	//fmt.Println(unicode.MaxRune)
	fmt.Println(unicode.IsSpace(' '))
	fmt.Printf("%#U", unicode.ToLower('G'))
}
