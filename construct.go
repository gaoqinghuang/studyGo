package main
import (
	"fmt"
)

type Cat struct {
	Color string
	Name  string
}

type BlackCat struct {
	Cat
}

func NewCat(name string) *Cat{
	return  &Cat{
		Name : name,
	}
}

func NewBlackCat(color string) *BlackCat{
	cat := &BlackCat{}
	cat.Color = color

	return cat
}

//实现构造函数
func main(){
	cat := NewBlackCat("black")
	fmt.Println(cat)
}