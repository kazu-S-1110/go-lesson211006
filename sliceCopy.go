package main

import "fmt"

func main() {
	/*
		sl := []int{100, 200}
		sl2 := sl

		sl2[0] = 1000
		fmt.Println(sl) //[1000 200]
		// スライスはデータ型なので代入をしても参照しているものは同じなことに注意
	*/
	// 上を回避するためにcopy関数を使用する

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl)  //copy(コピー先, コピー元) ここのnは変換に成功した数を入れる
	fmt.Println(n, sl2) //5 [1 2 3 4 5]

	// 別のアドレスを参照しているのか確認
	sl[2] = 400
	fmt.Println(sl)  //[1 2 400 4 5]
	fmt.Println(sl2) //[1 2 3 4 5]
}
