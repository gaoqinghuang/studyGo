package main
import (
	"fmt"
	"math/rand"
	"time"
)


//通道
func main() {
	// var channel = make(chan string)
	// go producer("cat",channel)
	// go producer("dog",channel)

	// customer(channel)

	//循环
	c := make(chan int)
	go func() {
	    c <- 1
	    c <- 2
	    c <- 3
	    close(c)
	}()
	for v := range c {
	    fmt.Println(v)
	}
}

func producer(header string,channel chan<- string) {
	
	for {
		channel<- fmt.Sprintf("%s:%v",header,rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel
		fmt.Println(message) 
	}
}

