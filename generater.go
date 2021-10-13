package main

import "fmt"

// ジェネレーター
// クロージャーの性質の利用すると実行する度に増分する関数が作れる

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()  //integersを再代入するとiはリセットされる
	fmt.Println(otherInts()) //1
}
