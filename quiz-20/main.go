package main

import (
	"fmt"
)

func main() {
	n := len("株式会社10ANTZ")
	if n == 10 {
		fmt.Println("X")
	} else {
		fmt.Println("Y")
	}
}
