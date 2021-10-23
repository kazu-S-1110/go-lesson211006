package main

import "fmt"

// struct 埋め込み
// 1つの構造体を、別の構造体に埋め込むこと

type T struct {
	// User User
	User //上の書き方を省略して書くことも可能
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	t := T{User: User{"user1", 10}}
	fmt.Println(t)           //{{user1 10}}
	fmt.Println(t.User)      //{user1 10}
	fmt.Println(t.User.Name) //user1
	// 省略した書き方のみ下のコードはアクセスが可能
	fmt.Println(t.Name) //user1

	// 埋め込んだ構造体の関数の渡し方
	t.User.SetName("hoge")
	fmt.Println(t.Name) //hogehoge
}
