package main

import "fmt"

func main() {
	var x = 12
	var max = 100

	if x > max {
		max = x
	}

	fmt.Println(max)

	var y = 1
	if x > y {
		fmt.Println("Big")
	} else if x < y {
		fmt.Println("Small")
	} else {
		fmt.Println("Equal")
	}

	var mode = "running"

	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}
}
