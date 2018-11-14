package main
import "fmt"

func main() {
	// var age = 24
	// if age >25{
	// 	fmt.Println("true")
	// }else if age == 25{
	// 	fmt.Println("==")
	// }else{
	// 	fmt.Println("<")
	// }

	// if err := "www";err != ""{//err变量只在if，else里面有效
	// 	fmt.Println(err)
	// 	return
	// }

	// var i = 0
	// for{
	// 	if i>10{
	// 		break
	// 	}
	// 	i++
	// }
	// fmt.Println(i)


	OuterLoop://break也有类似的效果，都是设置跳出的循环，本来只是跳出内层的
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            switch j {
            case 2:
                fmt.Println(i, j)
                continue OuterLoop
            }
        }
    }
	
	// var grade = "A"
	// switch {
	// 	case "C" == grade:
	// 		fmt.Println("A")
	// 	case "A" == grade:
	// 		fmt.Println("B")
	// 	case "A" == grade:
	// 		fmt.Println("C")
	//		fallthrough  //go的switch自带break，要去掉break的话加上fallthrough,不建议使用
	// 	default:
	// 		fmt.Println("D")
	// }
}