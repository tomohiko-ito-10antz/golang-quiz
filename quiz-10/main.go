package main

import (
	"fmt"
)

func f() {
	defer func() { fmt.Println("X") }()
	panic("fail")
}
func main() {
	f()
	fmt.Println("Y")
}
