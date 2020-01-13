package main

import (
	"fmt"
	"time"
)

func main() {
	//这个包很乱，非常乱，还要再回来看一次，主要是时区问题！！

	//c := time.Tick(1 * time.Second)
	//for now := range c {
	//	fmt.Printf("%v\n", now)
	//}

	//a := time.Now()
	//fmt.Println(time.Now().In(time.Local).Equal(time.Now()))

	//获取本地时间
	//t, _ := time.ParseInLocation("2006-01-02", "2019-06-24 14:26:16",time.Local)
	//fmt.Println(t.Zone())

	//a := time.Time{}
	//fmt.Println(a.String())

	//loc, _ := time.LoadLocation("UTC")
	//fmt.Println(loc.String())
	//local := time.Local
	//fmt.Println(local.String())

	//正确的，是中国时间
	//fmt.Println(time.Now().YearDay())

	//fmt.Println(time.Now().Unix())
	//time.Sleep(time.Second)
	//fmt.Println(time.Now().Unix())
}
