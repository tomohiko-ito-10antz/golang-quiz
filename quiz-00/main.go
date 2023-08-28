package main

import "fmt"

func f(a any) {
	if a == nil {
		fmt.Println("X")
	} else {
		fmt.Println("Y")
	}
}

func main() {
	var a *int = nil
	f(a)
}
