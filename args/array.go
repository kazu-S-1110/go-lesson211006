package main

import "fmt"

func main() {
	//  他の言語と違って要素数を変更できないので注意
	var arr1 [3]int          //何も中身を宣言しないパターン
	fmt.Println(arr1)        // [0 0 0] intの初期値の0が3つ入った。
	fmt.Printf("%T\n", arr1) //[3]int 型情報として3が格納されているため要素数を変更できないらしい

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) //[A B ]3つ目にはstringの初期値の空文字が入った。

	// 暗黙的な宣言方法
	arr3 := [3]string{"hoge", "sage", "fuga"}
	fmt.Println(arr3) //[hoge sage fuga]

	// 要素数を省略する方法（自動でカウントしてくれる
	arr4 := [...]string{"hoge", "sage", "foo", "fuga"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4) //[4]string

	// 配列の一部を参照する
	fmt.Println(arr4[0])   //hoge
	fmt.Println(arr4[0:3]) //0番目から3つ取り出す [hoge, sage, foo]

	// 要素の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// 要素数を調べる関数
	fmt.Println(len(arr1))
}
