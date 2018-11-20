package main
import (
	"fmt"
)

type MyInt int

func (m MyInt) IsZero() bool{
	return m == 0
}

func (m MyInt) Add(other MyInt) MyInt{
	return other + MyInt(m)
}

func main(){
	var b MyInt
	// b := 0
	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))
}