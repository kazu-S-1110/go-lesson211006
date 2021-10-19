package main

import "fmt"

// スライス
// 宣言、操作
func main() {
	// 空のスライスを定義
	var sl []int
	fmt.Println(sl) // []

	// 明示的宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // [100 200]

	// 暗黙的宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3) //[A B]

	// make関数を使用してsliceを作成
	sl4 := make([]int, 5)
	fmt.Println(sl4) //[0 0 0 0 0]

	// 値を変更する
	sl2[0] = 1000
	fmt.Println(sl2) //[1000 200]

	// 値の取り出し
	fmt.Println(sl2[0]) // 1000

	// 複数取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[1:]) //[2 3 4 5]
	fmt.Println(sl5[:3]) //[1 2 3]

	// 最後の要素を取り出す
	fmt.Println(sl5[len(sl5)-1]) //5
}
