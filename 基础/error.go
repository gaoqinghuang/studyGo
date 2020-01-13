package main

import (
	"fmt"
	// "os"
	// "errors"
)

/*
* 这个是一个不好的教程，看正常的看下面那个
 */
// type DivideError struct {
// 	divider int
// 	dividee int
// }

// func (de *DivideError) Error() string{
// // 	strFormat := `
// //     Cannot proceed, the divider is zero.
// //     dividee: %d
// //     divider: 0
// // `
// 	strFormat := `Cannot proceed, the
// 	divider is zero.dividee: %d divider: 0`  //""和··的区别在于一个会保留空格空行，一个不会
// 	return fmt.Sprintf(strFormat,de.dividee)
// }

// //这里的result errorMsg等于是不用定义
// func Divide(varDividee int,varDivider int)(result int,errorMsg string) {
// 	if varDivider == 0{
// 		dData := DivideError{
// 			dividee:varDividee,
// 			divider:varDivider,
// 		}
// 		errorMsg = dData.Error()
// 		fmt.Println(errorMsg)
// 		return //为空的话就是直接按方法定义好的直接返回
// 	}else{
// 		return varDividee / varDivider,""
// 	}
// }

// func main() {

// 	if result,errorMsg := Divide(100,10); errorMsg == ""{
// 		fmt.Println("100/10 = ",result)
// 	}//这里的result,errorMsg为局部变量，只在if里有用,else当然也能用

// 	if _,errorMsg := Divide(100,0); errorMsg != ""{
// 		fmt.Println("errorMsg is:",errorMsg)
// 	}
// }

/*
*看这个
 */

// var errDivisionByZero = errors.New("divsion by zero")

// func div(dividend,divisor int) (int,error){
// 	if divisor == 0{
// 		return 0,errDivisionByZero
// 	}

// 	return dividend/divisor,nil
// }

// func main() {
// 	fmt.Println(div(1,0))
// }

//自定义错误

type ParseError struct {
	Filename string
	Line     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

func newParseError(Filename string, line int) error {
	return &ParseError{Filename, line}
}

func main() {
	var e error
	e = newParseError("main.go", 1)

	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename:%s Line:%d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
