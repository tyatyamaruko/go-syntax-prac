package main

import (
	"fmt"
	"os"
)

func main() {

	// デフォルトフォーマット改行あり
	fmt.Println("hello golang")
	// 書き込み先を明示的に指定
	fmt.Fprintln(os.Stdout, "go lang")

	// フォーマットした結果を文字列で返す
	hello := fmt.Sprint("Hello Wolrd")
	fmt.Println(hello)

	fmt.Printf("%#v\n", hello)
}
