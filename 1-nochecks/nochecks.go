package main

import "fmt"

//test
func main() {
	fmt.Println("This program outputs a/b")
	fmt.Print("Enter a and b: ")
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println("a/b =", a/b)
}
