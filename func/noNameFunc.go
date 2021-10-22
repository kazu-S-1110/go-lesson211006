package main

import "fmt"

// 無名関数

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 3)
	fmt.Println(i)

	// 他のパターン
	i2 := func(x, y int) int {
		return x * y
	}(6, 5)
	fmt.Println(i2)

	f2 := ReturnFunc()
	f2() //I'm a function
}

// 関数を返す関数（コールバック関数やないかい）
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}
