package main

import "fmt"

func main() {
	var a any
	if _, ok := a.(any); ok {
		fmt.Println("X")
	} else {
		fmt.Println("Y")
	}
}
