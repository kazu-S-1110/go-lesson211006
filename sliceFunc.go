package main

import "fmt"

// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	//appendを使用して要素を追加
	sl = append(sl, 300)
	fmt.Println(sl)

	// 複数追加することも可能
	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	// makeはsliceを作成する関数
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	// lenは要素数を調べる
	fmt.Println(len(sl2))

	// capは入る容量を調べる
	// makeの第三引数に容量を設定できるが、パフォーマンスに影響するので今は意識しなくていい
}
