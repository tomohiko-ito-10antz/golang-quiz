package main

import (
	"fmt"
)

func f() string {
	b := "X"
	defer func() { b = "Y" }()
	return b
}
func main() {
	fmt.Println(f())
}
