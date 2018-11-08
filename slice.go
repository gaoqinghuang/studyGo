package main
import "fmt"

// var a = make([] int,3,5)
var a = [] int{}
func main() {
	// s :=[] int {1,2,3 } 
	printSlice(a)
	a = append(a,0)
	printSlice(a)
	a = append(a,1)
	printSlice(a)
	a= append(a,2,3,4,5,6)
	printSlice(a)
	var b = make([] int,len(a),cap(a)*2)
	printSlice(b)
	copy(b,a)
	printSlice(b)
	// fmt.Println(a,len(a),cap(a))
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}