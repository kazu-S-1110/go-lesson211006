package main

import "fmt"

// 関数

func Plus(x int, y int) int {
	return x + y
}

// 引数の型が同一の場合、省略することも可能
func multiply3(x, y, z int) int {
	return x * y * z
}

// 返り値が複数ある場合
func Divide(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値に変数を宣言するパターン
func Double(price int) (result int) {
	result = price * 2
	return
}

// 返り値が無いパターン
func NoReturn() {
	fmt.Println("No return")
}

func main() {
	fmt.Println(Plus(1, 4))
	fmt.Println(multiply3(3, 4, 5))
	q, r := Divide(75, 4)
	fmt.Println(q, r) //18  3
	// 返り値の破棄（使われてない変数はエラーを吐くが_にすることで破棄することが可能）
	q2, _ := Divide(53, 4)
	fmt.Println(q2)
	NoReturn()
}
