package main

import "fmt"

func main() {
	byteA := []byte{72, 73} //配列の宣言方法、配列にbyte型で72,73を代入している感じ
	fmt.Println(byteA)      //[72 73]

	//文字列へ変換
	fmt.Println(string(byteA)) //HI

	// 文字列をbyteの配列（スライスって言い方？）に変換する方法
	c := []byte("HI")
	fmt.Println(c) //[72 73]

	//これ使い道あるん？？？
}
