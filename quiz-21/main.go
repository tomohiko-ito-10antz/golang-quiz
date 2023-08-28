package main

import (
	"fmt"
)

func main() {
	for i := range "いろは" {
		if i == 2 {
			fmt.Println("X")
		}
		if i == 6 {
			fmt.Println("Y")
		}
	}
}
