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
}
