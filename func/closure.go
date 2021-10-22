package main

import "fmt"

// クロージャー
// クロージャを使うと、関数内で状態を保持しつつ処理を行うことができる。
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))

}
