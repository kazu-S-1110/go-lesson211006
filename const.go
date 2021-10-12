package main

import "fmt"

// 定数の扱い方
// 定数はグローバルに宣言することが多い
const Pi = 3.14 //定数名の頭を大文字にすることで他のパッケージからも呼び出せる。
// 裏を返せば大文字以外ならこのパッケージ内でしか使えない

// 複数宣言する場合
const (
	URL      = "http://google.com"
	SiteName = "test"
)

// 値を省略して宣言する場合
const (
	A = 1
	B
	C
	D = "HELLO"
	E
	F
)

// 大きな値の定義も可能(ただし使用することは不可、varでの定義は無理)
var Big int = 9223372036854775807 //限界値
const Big2 = 9223372036854775807 + 1

// iota（連番を作成する）
const (
	c0 = iota
	c1
	c2
	c3
)

func main() {
	fmt.Println(Pi)
	// Pi = 3.4 //cannot assign to Pi (declared const)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F) //1 1 1 HELLO HELLO HELLO //省略して代入することが可能

	fmt.Println(Big2 - 1)

	fmt.Println(c0, c1, c2, c3) //0 1 2 3
}
