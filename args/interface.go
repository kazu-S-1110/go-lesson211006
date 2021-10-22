package main

import "fmt"

func main() {
	// interfaceについて学習 全ての型と互換性があり、nilも入れるということだけまず覚える
	var x interface{}
	fmt.Println(x) //nil 何も宣言していないのでnilとなる

	// interfaceはどんな値も代入することができる
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "hoge"
	fmt.Println(x)
	x = [4]int{1, 3, 5}
	fmt.Println(x)

	// ただしデータ特有の演算などは不可なので注意。
	x = 1
	// fmt.Println(x + 4) //invalid operation: mismatched types interface{} and int

}
