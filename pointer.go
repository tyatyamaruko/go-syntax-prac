package main

import "fmt"

func main() {
	var a1 int
	var p1 *int

	p1 = &a1
	*p1 = 123

	fmt.Println(a1)
}
