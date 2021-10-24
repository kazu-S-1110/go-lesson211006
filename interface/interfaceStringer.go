package main

import "fmt"

// interface
// fmt.Stringer

// 中身
// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

// 下のメソッドを定義した場合、文字列として出力することができる
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	// 下の一行は文字列として出力しない場合。ポインタとして認識される
	// fmt.Println(p) //&{100 ABC}
	// 下の一行はString()を定義した場合。
	fmt.Println(p) //<<100, ABC>> しっかりstring（文字列）として出力されている
}
