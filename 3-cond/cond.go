package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program outputs a/b")
	fmt.Print("Enter a and b: ")
	var a, b float64
	fmt.Scan(&a, &b)
	for b == 0 { // Loops while b == 0.
		fmt.Println("ERROR: cannot divide by zero")
		fmt.Print("Enter a and b: ")
		fmt.Scan(&a, &b)
	}
	fmt.Println("a/b =", a/b)
}
