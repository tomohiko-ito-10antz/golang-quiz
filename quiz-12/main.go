package main

import (
	"fmt"
)

func f() (x string) {
	b := "X"
	defer func() { b = "Y" }()
	return b
}
func main() {
	fmt.Println(f())
}
