package main
import "fmt"

//实现阶乘
func main() {
	i := 10
	fmt.Println(Factorial(i)) 
}

func Factorial(n int)(result int){
	if n>0{
		result = n*Factorial(n-1)
		return result
	}
	return 1
}