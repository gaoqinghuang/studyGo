package main
import "fmt"


type Phone interface{
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
}