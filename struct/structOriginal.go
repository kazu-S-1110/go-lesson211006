package main

import "fmt"

// structの独自型
// 既存の型をオーバーライドする
// 既存の型と演算ができないのでそれを利用して活用される。

type MyInt int

// 独自型にメソッドを定義
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi) //main.MyInt

	// i := 100
	// fmt.Println(i + mi) //error invalid operation: mismatched types int and MyInt

	mi.Print()
}
