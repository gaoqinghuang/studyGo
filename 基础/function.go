package main

import "fmt"

//给结构体定义方法
// type Circle struct{
// 	radius float64
// }

// func main() {
// 	var c Circle
// 	c.radius = 1
// 	fmt.Println("这个圆的面积是：",c.getArea())
// }

// func (c Circle) getArea() float64 {
// 	return 3.14*c.radius*c.radius
// }

func main() {
	// var a string = "12"
	// var a int = 10
	// var b int = 12
	// var c string = "12"
	// var res int
	// res = max(a,b,c)
	// fmt.Println(res)

	// fmt.Println(namedRetValues())

	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	fmt.Printf("%+v\n", in)
	fmt.Printf("%+p\n", &in)

	out := passByValue(in)

	fmt.Printf("%+v\n", out)
	fmt.Printf("%+p\n", &out)
}

func max(num1, num2 int, num3 string) int {
	// var res int
	// if num1>num2{
	// 	res = num1
	// }else{
	// 	res = num2
	// }
	return 1
}
func namedRetValues() (a, b int) {
	a = 1
	return a, b
}

//参数的传递
type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("%+v\n", inFunc)
	fmt.Printf("%p\n", &inFunc)

	return inFunc
}
