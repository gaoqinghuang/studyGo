package main

import (
	"encoding/json"
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

var struct_pointer *Books

//匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

func main() {
	//book := Books{"解忧杂货","小高"}
	book := Books{title: "解忧杂货", author: "小高"}
	data, _ := json.Marshal(book)
	fmt.Println(string(data))
	//var abb Books
	//fmt.Println(json.Unmarshal(data,&abb))
	//fmt.Println(book)
	// struct_pointer = &book

	// // book.title = "w"
	// fmt.Println(struct_pointer.title)

	// var book Books
	// book := new(Books)
	// book := &Books{}
	// book.title = "解忧杂货"
	// fmt.Println(book)

	// msg := &struct{
	// 	id int
	// 	data string
	// }{
	// 	1024,
	// 	"hello",
	// }
	// printMsgType(msg)
}
