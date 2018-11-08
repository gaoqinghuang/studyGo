package main
import "fmt"

//给结构体定义方法
type Circle struct{
	radius float64
}

func main() {
	var c Circle
	c.radius = 1
	fmt.Println("这个圆的面积是：",c.getArea())
}

func (c Circle) getArea() float64 {
	return 3.14*c.radius*c.radius
}

// func main() {
// 	// var a string = "12"
// 	var a int = 10
// 	var b int = 12
// 	var c string = "12"
// 	var res int
// 	res = max(a,b)
// 	fmt.Println(res)
// }

// func max(num1,num2 int,var num3 string = 'w') int{
// 	// var res int
// 	// if num1>num2{
// 	// 	res = num1
// 	// }else{
// 	// 	res = num2
// 	// }
// 	return 1
// }

