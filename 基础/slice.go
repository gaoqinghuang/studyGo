package main

import "fmt"

// var a = make([] int,3,5)
var a = []int{}
var b = [3]int{1, 2, 3}

func main() {
	// s :=[] int {1,2,3 }
	// printSlice(a)
	// a = append(a,0)
	// printSlice(a)
	// a = append(a,1)
	// printSlice(a)
	// a= append(a,2,3,4,5,6)
	// printSlice(a)
	// var b = make([] int,len(a),cap(a)*2)
	// printSlice(b)
	// copy(b,a)
	// printSlice(b)
	// fmt.Println(a,len(a),cap(a))

	//数组变切片
	// fmt.Println(b[1:])
	// for k,v := range b[1:]{
	// 	fmt.Println(k,v)
	// }
	//生成空切片
	// fmt.Println(b[0:0])
	// var c [] string
	// var d = [] string{}
	// fmt.Println(c,d)
	// fmt.Println(c == nil)//true 只是声明
	// fmt.Println(d == nil)//false 赋值了，只是赋的空值

	//动态的创建一个切片
	// e := make([]int,2,3)
	// fmt.Println(e == nil)

	//添加元素
	// var f []int
	// for i := 0; i < 10; i++ {
	// 	f = append(f,i)
	// 	printSlice(f)
	// }
	// var g []string
	// g = append(g,"xiao")
	// g = append(g,"gao","xiao")
	// h := []string{"gao","xiao"}
	// g = append(g,h...)//这个的...不能丢
	// fmt.Println(g)

	//复制
	// var h =make([] int,2)
	// h = append(h,4,5,6,7)
	// printSlice(h)
	// var i = [] int{1,2,3,4}
	// copy(h,i)
	// printSlice(h)

	// var a = make([]int,3)
	// for i := 0; i < 3; i++ {
	// 	a[i] = i
	// }

	// var a = [3]int{0,1,2}
	// c  := a[1:3]
	// b := c//切片这样赋值改变c时b也会跟着变，如果是数组则不会
	// fmt.Println(c)
	// c[0] = 10
	// fmt.Println(b)

	//删除元素
	seq := []string{"a", "b", "c", "d", "e"} //切片,如果[]里面有带数字就是数组

	seq = append(seq[0:2], seq[3:5]...)
	fmt.Println(seq)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
