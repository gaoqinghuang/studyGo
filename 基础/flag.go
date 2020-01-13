package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode1", "1", "process mode")

func main() {
	flag.Parse()
	fmt.Println(*mode)
}
