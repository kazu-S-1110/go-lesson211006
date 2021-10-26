package main

import "fmt"

// fmt

func main() {
	fmt.Println("display")

	// fmt標準
	fmt.Print("Hello") //改行なし
	// 改行あり
	fmt.Println("Hello")

	// Println系　可変長引数おk。　各の文字列は半角スペースで区切られる。文字列の最後に改行が加わる。
	fmt.Println("hello", "world", "!", "yeeee") //hello world ! yeeee

	// Printf系　書式指定子によって表示するものが変わってくる。記事見ながらやりましょhttps://qiita.com/rock619/items/14eb2b32f189514b5c3c
	// 論理値(bool)	%t
	// 符号付き整数(int, int8など)	%d
	// 符号なし整数(uint, uint8など)	%d
	// 浮動小数点数(float64など)	%g
	// 複素数(complex128など)	%g
	// 文字列(string)	%s
	// チャネル(chan)	%p
	// ポインタ(pointer)	%p}
	fmt.Printf("%s\n", "hello")  //hello
	fmt.Printf("%#v\n", "Hello") //"Hello"
}
