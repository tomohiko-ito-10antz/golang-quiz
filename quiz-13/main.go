package main

import (
	"fmt"
)

func main() {
	c := "X"
	defer fmt.Println(c) // A
	c = "Y"
	defer fmt.Println(c) // B
	c = "Z"
}
