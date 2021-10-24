package main

import "fmt"

//interface
//interfaceは単純な型（なんでも入ってくれる型）と
// 関数をまとめる型を持っている。
// structは異なるけどmethodは同じような場合、それぞれにメソッドを定義するのは冗長なので
// interfaceでまとめてあげると便利。
// https://medium.com/since-i-want-to-start-blog-that-looks-like-men-do/初心者に送りたいinterfaceの使い方-golang-48eba361c3b4

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

// 二つの共通をまとめる型を定義
type Stringfy interface {
	ToString() string
}

func main() {
	// まとめた型のスライスを宣言。
	vs := []Stringfy{
		&Person{"hoge", 32},
		&Car{"32-43", "AB-3243"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
		// 実行結果
		// Name=hoge, Age=32
		// Number=32-43, Model=AB-3243
	}
}
