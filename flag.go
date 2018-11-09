package main
import (
	"fmt"
	"flag"
)

var mode = flag.String("mode1","1","process mode")

func main() {
	flag.Parse()
	fmt.Println(*mode)
}