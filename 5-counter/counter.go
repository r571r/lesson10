package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++ // x += 1 or x = x + 1
	}
}
