package main

import "fmt"

func main() {
	var a any
	switch a.(type) {
	case any:
		fmt.Println("X")
	default:
		fmt.Println("Y")
	}
}
