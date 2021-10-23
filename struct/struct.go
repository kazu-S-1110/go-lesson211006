package main

import "fmt"

// struct

type User struct {
	Name string
	Age  int
	// X, Y int  複数指定することも可能
}

func UpdateUser(user User) {
	user.Name = "hoge"
	user.Age = 1000
}

func UpdateUser2(user *User) { //structは参照渡しで値を更新するべし！
	user.Name = "foo"
	user.Age = 999
}

func main() {
	// 何も指定しないとデフォルト値が入る
	var user1 User
	fmt.Println(user1) // { 0}
	user1.Name = "user1"
	user1.Age = 49
	fmt.Println(user1)

	// 暗黙的宣言
	user2 := User{}
	fmt.Println(user2)

	// 初期値を最初から指定
	user3 := User{Name: "user3", Age: 39}
	fmt.Println(user3)

	// 省略形
	user4 := User{"user4", 20} //順番（field）は守って宣言する
	fmt.Println(user4)

	//一部だけ宣言することも可能
	user6 := User{Name: "user6"}
	fmt.Println(user6)

	// user7 := User{"user7"} //一部だけを宣言する場合は省略形はだめ
	// fmt.Println(user7)

	// newを使うとポインタを返す
	user8 := new(User)
	fmt.Println(user8) //&{ 0}

	// 別の書き方（こっちの方が多いらしい
	user9 := &User{}
	fmt.Println(user9)

	// 関数で値を更新するときは参照渡しで行う
	UpdateUser(user1)
	UpdateUser2(user9)
	fmt.Println(user1) //{user1 49}
	fmt.Println(user9) //&{foo 999}
}
