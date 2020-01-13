package main

import "fmt"

func main() {
	// var a = []int{1,2,3,4}
	// // var sum int
	// // sum :int
	// for i,num := range a{
	// fmt.Println(i,"=>",num)
	// }
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
