package main
import "fmt"
// import "unsafe"

const (
	a = iota
	b 
	c 
	d = "ha"
	e 
	f = 100
	g 
	h = iota
	i 
)
var z = 10
var j *int
func main() {
	j = &z
	// var s
	fmt.Println("%T",z)
	fmt.Println("%d",j)
	fmt.Println("%d",&z)
	// fmt.Println(a,b,c,d,e,f,g,h,i)
		// fmt.Println(a,b,c)
}