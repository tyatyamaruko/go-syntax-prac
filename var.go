package main

import "fmt"

func main() {
	// 変数宣言
	var a1 int = 123
	// 型推論
	var a2 = 123
	// 宣言省略形
	a3 := 123
	// まとめて宣言
	var (
		a4 = 123
		a5 = 456
	)
	name, age := "Yamada", 26

	const foo = 100
	const (
		hogehoge = 100
		fugafuga = 200
	)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(hogehoge)
	fmt.Println(fugafuga)
}
