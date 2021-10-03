package main

import "fmt"

func main() {
	fmt.Println("This program verifies whether a number is a palindrome.")
	fmt.Print("Enter a number: ")
	var x uint // Only positive integral numbers.
	fmt.Scan(&x)
	var y uint
	tmp := x
	for tmp != 0 {
		y = y*10 + tmp%10
		tmp /= 10
	}
	if x == y {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
