package main

import "fmt"

/*type Phone interface{
	call()
}

type NokiaPhone struct{
}

func (nokiaPhone NokiaPhone) call(){
	fmt.Println("nokia")
}

type Iphone struct{

}

func (iphone Iphone) call() {
	fmt.Println("Iphone")
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
}*/

type DataWriter interface {
	// WriteData(data interface{}) error
	CanWrite() bool
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}
func (d *file) CanWrite() bool {
	fmt.Println("CanWrite:")
	return true
}

func main() {
	f := new(file)

	var writer DataWriter
	writer = f

	writer.CanWrite()
}
