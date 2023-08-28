package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("X")
	os.Exit(1)
}
