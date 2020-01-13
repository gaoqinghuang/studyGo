package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Local)
	fmt.Println(time.LoadLocation("Local"))
	fmt.Println(time.Now())
	//fmt.Println(time.Parse("2006-01-02 15:04:05", "2018-04-23 12:24:51"))
	acc, _ := time.Parse("2006-01-02 15:04:05", "2017-05-11 14:06:06")
	fmt.Println(acc.Format("2006-01-02 15:04:05"), acc.Unix())
	//fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local))
	abb, _ := time.LoadLocation("Local")
	a, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", abb)
	fmt.Println(a.Format("2006-01-02 15:04:05"), a.Unix())
}
