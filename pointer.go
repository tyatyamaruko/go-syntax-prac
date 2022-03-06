package main

import "fmt"

func main() {
	var a1 int = 123
	var a2 int = 123
	// var p1 *int

	// p1 = &a1
	// *p1 = 123
	fn(a1, &a2)
	fmt.Println(a1, a2)
}

func fn(b1 int, b2 *int) {
	b1 = 456
	*b2 = 456
}
