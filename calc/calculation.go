package main

import (
	"fmt"
	"math"
)

// 算術演算子
func main() {
	fmt.Println(1 + 2)   //3
	fmt.Println(4 - 2)   //2
	fmt.Println(99 * 4)  //396
	fmt.Println(945 / 7) //135

	// 余を求める計算
	fmt.Println(654 % 9) //6

	// 冪乗は外部モジュールを使用する
	fmt.Println(math.Pow(2, 4)) //16

	// 文字列をつなげる
	fmt.Println("ABCDE" + "hogehoge") //ABCDEhogehoge

	// 算術演算子と代入
	n := 0
	n += 3
	fmt.Println(n) //3
	n++
	n -= 4
	n--
	fmt.Println(n) //-1
	n *= 3
	fmt.Println(n) //-3
	n /= 3
	fmt.Println(n) //-1

	// 文字列にも可能
	s := "Hello"
	s += "World"
	fmt.Println(s) //HelloWorld
}
