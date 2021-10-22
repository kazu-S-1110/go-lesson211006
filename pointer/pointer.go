package main

import "fmt"

// pointer

func Double(i int) {
	i = i * 2 ///nを渡しても元のnは変更されず、新しいnがコピーされて作られている。
}
func Doublev2(i *int) { //参照渡しと呼ばれる手法
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n) //0xc0000b2008 この変数のメモリ上のアドレスを表示している

	Double(n)
	fmt.Println(n) //100

	// ポインタ変数を宣言
	var p *int = &n
	fmt.Println(p) //0xc0000b2008
	// 参照先を表示するには以下
	fmt.Println(*p) //100

	// 元の値を変更するには以下
	*p = 300
	fmt.Println(n)  //300
	fmt.Println(*p) //300
	n = 5000
	fmt.Println(n)  //5000
	fmt.Println(*p) //5000

	//関数もポインタを利用して使用する
	Doublev2(&n)
	fmt.Println(n)  //10000
	fmt.Println(*p) //10000

	Doublev2(p)
	fmt.Println(n)  //20000
	fmt.Println(*p) //20000

}
