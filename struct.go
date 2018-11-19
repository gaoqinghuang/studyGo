package main
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

var struct_pointer *Books

//匿名结构体
func printMsgType(msg *struct{
	id int
	data string
}){
	fmt.Printf("%T\n",msg)
}

func main() {
	// var book = Books{"解忧杂货","小高","解密",12}
	// struct_pointer = &book

	// // book.title = "w"
	// fmt.Println(struct_pointer.title)

	// var book Books
	// book := new(Books)
	// book := &Books{}
	// book.title = "解忧杂货"
	// fmt.Println(book)

	msg := &struct{
		id int
		data string
	}{
		1024,
		"hello",
	}
	printMsgType(msg)
}