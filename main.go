package main

import "fmt"

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 1000
	fmt.Println(i)

	var word string = "Hello, Go"
	fmt.Println(word)

	// まとめて定義する
	var t, f bool = true, false
	fmt.Println(t, f)

	// 異なる型だけどまとめて定義する
	var (
		i2    int    = 2000
		word2 string = "Good afternoon, Go"
	)
	fmt.Println(i2, word2)
}
