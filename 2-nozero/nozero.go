package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program outputs a/b")
	fmt.Print("Enter a and b: ")
	var a, b float64
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Println("ERROR: cannot divide by zero")
		return
	}
	fmt.Println("a/b =", a/b)
}
