package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 真偽値を文字列に変換するsd   600
	bt := true
	fmt.Println(strconv.FormatBool(bt))        //true
	fmt.Printf("%T\n", strconv.FormatBool(bt)) //string

	// 整数を文字列に変換する
	i := strconv.FormatInt(-100, 10) //第二引数は進数を指定可能
	fmt.Printf("%v, %T\n", i, i)     //-100, string

	// 簡易的に変換できる
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2) //100, string

	// 浮動小数点型を文字列に
	// 第二引数は書式指定子、使う時に調べて使おう
	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))

	// 文字列を真偽値に変換する
	// trueへ変換できる文字列
	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", bt1, bt1)
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6) //みんなtrue

	// falseへ変換できる文字列
	// 二番目の戻り値はerror型、エラーハンドリングが可能
	bf1, ok := strconv.ParseBool("false")
	if ok != nil {
		fmt.Println("Convert Error")
	}
	fmt.Printf("%v, %T\n", bf1, bf1) //false, bool
	bf2, _ := strconv.ParseBool("0")
	bf3, _ := strconv.ParseBool("f")
	bf4, _ := strconv.ParseBool("F")
	bf5, _ := strconv.ParseBool("FALSE")
	bf6, _ := strconv.ParseBool("False")
	fmt.Println(bf2, bf3, bf4, bf5, bf6) //false false false false false

}
