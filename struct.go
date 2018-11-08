package main
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

var struct_pointer *Books
func main() {
	var book = Books{"解忧杂货","小高","解密",12}
	struct_pointer = &book

	// book.title = "w"
	fmt.Println(struct_pointer.title)
}