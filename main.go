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

	//変数だけを宣言
	var i3 int       //何も値を指定しない場合、intの初期値は0が入る
	var word3 string //stringは空文字が入る

	fmt.Println(i3, word3)
	// 再代入
	i3 = 500
	word3 = "hoge"
	fmt.Println(i3, word3)

	// 暗黙的な定義方法 (型推論が働く)
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	// 再代入は可能
	i4 = 450
	// 再定義は不可
	// i4 := 500
	// また再代入の際、別の型を代入することもだめ
	// i4 = "hello"

	outer() //関数外の関数を呼び出すことは可能
	// fmt.Println(word4) //関数外の変数を呼び出すことは不可
	// var i6 int = 500 //使用してない変数にもエラーを吐いてくれる
	// ただし、関数外（グローバル）に宣言して変数はエラーを吐かない？
}

// また関数外で暗黙的な宣言はできない(明示的宣言は可能)
// i5 := 400

// 明示的宣言と暗黙的宣言の使い分けは関数外で使うか関数内で使うか大きいが
// 基本的には明示的に宣言すると良い。静的型付けの利点を活かすようにする

func outer() {
	var word4 string = "this is outer variable"
	fmt.Println(word4)
}
