package main

import "fmt"

// struct method
// 構造体を渡す関数は基本的に参照渡しでする

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name //これだと値渡しになって元の値は更新されない。参照渡しする必要あり
}

func (u *User) SetName2(name string) {
	u.Name = name //参照渡しをして元の値を更新させる
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName() //こういうかんじで使う

	user1.SetName("hoge")
	user1.SayName() //user1

	user1.SetName2("foo")
	user1.SayName() // foo

	user2 := &User{Name: "user2"}
	user2.SetName("sage") //ポインタ型で定義したものを値渡しで変更しても元の値は変わらない
	user2.SayName()       //user2

	user2.SetName2("sage")
	user2.SayName() //sage

	// 結論：構造体を参照型で作ろうがデータ型で作ろうが関数が参照渡しでないと元の値が変わらないので
	// 関数は参照渡しで作ろう。
}
