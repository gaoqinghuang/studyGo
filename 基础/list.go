package main
import (
	"fmt"
	"container/list"
)

func main() {
	// var a list.List
	// var b = list.New()
	// fmt.Println(a,b)
	l := list.New()
	l.PushBack("canon")
	l.PushFront(67)

	element := l.PushBack("fist")

	l.InsertAfter("high",element)
	l.InsertBefore("noon",element) //67, canon, noon, fist, high
	l.Remove(element) //	67, canon, noon, high

	for i := l.Front(); i !=nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}