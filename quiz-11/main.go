package main

import (
	"fmt"
)

func f() (a string) {
	defer func() { a = "X" }()
	return "Y"
}
func main() {
	fmt.Println(f())
}
