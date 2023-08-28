package main

import "fmt"

func main() {
	var a any
	switch a.(type) {
	default:
		panic("fail")
	case nil:
		fmt.Println("X")
	case any:
		fmt.Println("Y")
	}
}
