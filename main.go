package main

import (
	"fmt"

	"github.com/wailman24/git-workshop/functions"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("hello world !!")

	fmt.Println(functions.Multiply(1, 2))
	fmt.Println(functions.Subtract(5, 3))
	fmt.Println(add(3, 4))
}
