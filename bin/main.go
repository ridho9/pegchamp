package main

import (
	"fmt"

	"github.com/ridho9/pegchamp"
)

func main() {
	input := "hello world"
	p := pegchamp.String("hello")

	fmt.Println(p.Run(input))
}
