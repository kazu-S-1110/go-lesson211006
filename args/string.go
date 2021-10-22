package main

import "fmt"

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si) //数値を代入してもstring型のまま

	// 複数行の文字列
	fmt.Println(`test
			test
					test`)
	// 特殊文字の表示方法
	fmt.Println("\"") //バックスラッシュを用いる
	fmt.Println(`"`)  //全体をバッククォートで囲む

	// 文字列の操作
	fmt.Println(s[0])         //72が出力される string型はバイト型なので返ってくる値は数値になる メモリには文字列＝数値として扱われるらしい
	fmt.Println(string(s[0])) //Hが出力される

}
