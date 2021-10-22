package main

import "fmt"

func main() {
	var i int = 100
	/* intのみ場合環境依存になる,64ビットならint64
	int8        the set of all signed  8-bit integers (-128 to 127)
	int16       the set of all signed 16-bit integers (-32768 to 32767)
	int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	*/

	fmt.Println(i + 40)
	var i2 int64 = 200
	// 型が違う場合、エラーになる
	// fmt.Println(i + i2) 型が違うのでエラー

	// 型情報を調べる関数
	fmt.Printf("%T\n", i2) //int64

	// 型を変換する方法（int32とかで囲むだけ）
	fmt.Printf("%T\n", int32(i2))
	// 型が統一されたので演算が可能になる
	fmt.Println(i + int(i2))
}
