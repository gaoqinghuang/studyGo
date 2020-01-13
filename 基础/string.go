package main

import (
	"fmt"
	// "unicode/utf8"
	// "strings"
	"bytes"
)

func main() {
	// //打印长度
	// s1 := "我是谁"
	// s2 := "xiaogao"
	// fmt.Println(len(s1),len(s2),utf8.RuneCountInString(s1),utf8.RuneCountInString(s2))

	// //遍历
	// theme := "xiaogao小高"
	// for x,s := range theme {
	// 	fmt.Printf("%d:%c=>%d \n",x,s,s)
	// }

	// //截取
	// tracer := "死神来了，死神bye bye"
	// fmt.Println(tracer[12:15])
	// comma := strings.Index(tracer,"，")
	// againComma := strings.Index(tracer[comma:],"死神")
	// fmt.Println(comma,againComma)
	// fmt.Println(tracer[againComma+comma:1])

	// //修改
	// angel := "heros nerver die"
	// angelBytes := []byte(angel)
	// for i := 5; i < 10; i++ {
	// 	angelBytes[i] = ' '
	// }
	// fmt.Println(string(angelBytes))

	// //连接
	// //直观的+号连接法
	// start := "xiao"
	// end := "gao"
	// fmt.Println(start+end)
	// //高效的StringBuilder
	// var stringBuilder bytes.Buffer
	// stringBuilder.WriteString(start)
	// stringBuilder.WriteString(end)
	// fmt.Println(stringBuilder.String())

}
