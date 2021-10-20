package main

import "fmt"

// map
// JavaScriptでいうobject, key:valueの形
func main() {
	// 明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m) //map[A:100 B:200]

	// 暗黙的宣言（違いがほぼないw
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	// 改行してもおk
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// makeを使用して空のmapを作成
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	fmt.Println(m4)
	m4[4545] = "YEEEEE"
	fmt.Println(m4)

	// 取り出し方法
	fmt.Println(m2["A"])

	// 存在しないキーを指定してもエラーにはならず空文字が帰るのでエラーハンドリングをする
	value, boolean := m4[1]     //一つ目に取り出したvalue,二つ目に成功したかの判定が代入される
	fmt.Println(value, boolean) //JAPAN true
	value2, boolean2 := m4[15464]
	fmt.Println(value2, boolean2) // false
	if !boolean2 {
		fmt.Println("キー間違えてんがな〜〜〜")
	}

	// 要素の削除方法
	delete(m4, 1) //第一引数にmap, 第二引数にキーを指定
	fmt.Println(m4)
}
