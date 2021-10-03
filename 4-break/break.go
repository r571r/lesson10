package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program outputs a/b")
	var a, b float64 // Cannot be inside of for to be visible outside.
	for {            // Loops indefinitely.
		fmt.Print("Enter a and b: ")
		fmt.Scan(&a, &b)
		if b != 0 {
			break // Stops the loop.
		}
		fmt.Println("ERROR: cannot divide by zero")
	}
	fmt.Println("a/b =", a/b)
}
