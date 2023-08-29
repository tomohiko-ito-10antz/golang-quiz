package main

import (
	"fmt"
)

func main() {
	c := "Z"
	defer fmt.Println(c)
	c = "Y"
	defer fmt.Println(c)
	c = "X"
}
