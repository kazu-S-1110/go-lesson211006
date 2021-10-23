package main

import "fmt"

// 構造体とスライス

type User struct {
	Name string
	Age  int
}

// 構造体を入れるためのスライスを定義
type Users []*User

// 上のコードの別の書き方。冗長なので上の方が推奨
// type Users struce {
// 	Users []*Users
// }

func main() {
	user1 := User{"user1", 10}
	user2 := User{"user2", 20}
	user3 := User{"user3", 30}
	user4 := User{"user4", 40}

	users := Users{}

	users = append(users, &user1, &user2, &user3, &user4)
	fmt.Println(users) //[0xc00000c030 0xc00000c048 0xc00000c060 0xc00000c078]

	for _, v := range users {
		fmt.Println(*v)
	}

}
