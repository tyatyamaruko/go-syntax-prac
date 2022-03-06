package main

import "fmt"

func main() {
	fmt.Println(add(12, 3))

	funcA(1, 2, 3, 4, 5)
}

func add(x int, y int) int {
	return x + y
}

func funcA(a int, b ...int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d]=%d\n", i, num)
	}
}
