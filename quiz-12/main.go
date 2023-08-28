package main

import (
	"fmt"
)

func f() (x string) {
	b := "Y"
	defer func() { b = "X" }()
	return b
}
func main() {
	fmt.Println(f())
}
