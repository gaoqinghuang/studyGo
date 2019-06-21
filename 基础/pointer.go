package main
import "fmt"

//变量的声明
var a  = []int{2,3,4}
// var a  = []int{}

var ptr = [3]*int{}
 // var ptr []*int;


//指针
func main() {

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}
	for j := 0; j < len(a); j++ {
		fmt.Println(*ptr[j])
	}
	// fmt.Println(a)
	// ip = &a
	// fmt.Println(ip,&a)
	// fmt.Println(*ip)
	// fmt.Printf("%x",ip)
	// // if ip == nil{
	// 	fmt.Println("yes")
	// }else{
	// 	fmt.Println("no")
	// }

	//一个数组指针分别指向地址
}