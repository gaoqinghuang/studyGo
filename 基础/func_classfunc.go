package main
import (
	"fmt"
)

type class struct {
	
}

func (c class) Do(v int) {
	fmt.Println("class method",v)
}

func Do(v int){
	fmt.Println("function",v)
}

func main() {
	var delegate func(int)

	c := new(class)

	delegate = c.Do

	delegate(100)

	delegate = Do
    
    delegate(100)
}