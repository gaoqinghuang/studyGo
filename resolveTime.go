package main
import (
	"fmt"
)

func main() {
	fmt.Println(resolveTime(1000))
}

func resolveTime(s int)(d int,h int,m int) {
	d = s/(60*60*24)
	h = s/(60*60)
	m = s/(60)
	return
}

