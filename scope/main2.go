package main

import (
	// f "fmt" //パッケージ名の短縮（あまり推奨されない）
	"fmt"
	foo "foo/mypkg"
)

func appName() string {
	const AppName = "GoApp"
	var Version = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s //返り値にbを指定しているので重複しているためエラーを吐く
	b = s //これが正解

	// 関数内でもスコープは定義可能
	{
		b := "BBBB"
		fmt.Println(b) //BBBB
	}
	fmt.Println(b) // AAAA
	return b
}

func main() {
	fmt.Println(foo.Max)
	// fmt.Println(foo.min) //min not exported by package foo
	fmt.Println(foo.ReturnMin()) //1

	fmt.Println(appName()) //GoApp 1.0
	// fmt.Println(AppName, Version)   //undefined: AppName
	// 呼び出そうとしている定数や変数が小文字の変数内にあるためスコープが狭くなっている。

	fmt.Println(Do("AAAA")) //AAAA
}
