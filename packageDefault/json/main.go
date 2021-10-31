package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// json
// 構造体からjsonテキストへの変換

type A struct{}

type User struct {
	ID int `json:"id"` //json:"id,string"　とすると表示する際、stringになる
	// またjson:"-"とすると表示する際に非表示とすることが可能
	Name string `json:"name"`
	// もし、受け取る値がなく、それを表示させたくない場合は,json:"name,omitempty"とする
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal Json に変換 (Marshalの意味は組織化するとか構造化するとか)
	bs, err := json.Marshal(u) //byteのスライスで受け取る
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs)) //{"id":1,"name":"test","email":"test@test.com","created":"2021-10-31T19:33:05.546494+09:00","A":{}}

	// Jsonから構造体への変換
	fmt.Printf("%T\n", bs) //[]uint8

	u2 := new(User)

	// Unmarshal Jsonをデータに変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)        //&{1 test test@test.com 2021-10-31 19:41:45.085824 +0900 JST {}}
	fmt.Printf("%T\n", u2) //*main.User
	fmt.Println(u2.Name)
}
