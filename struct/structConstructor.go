package main

import "fmt"

// structのconstructor
// constructor関数はないけど同じようなものを作ることは多いらしい

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 34)
	// これで参照型の構造体が作れる
	fmt.Println(user1) //&{user1 34}
	// 実体にアクセスするなら以下
	fmt.Println(*user1) //{user1 34}

}
