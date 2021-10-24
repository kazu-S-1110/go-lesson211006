package main

import "fmt"

// interface
// カスタムエラー

// 標準のエラーは以下のようなコードで定義されているのでinterfaceで上書きして独自にカスタマイズできる

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

// 下のコードでErrorの構造を持った状態でカスタムされたErrorのメソッドを定義できる。
func (e *MyError) Error() string {
	return e.Message
}

// メソッドを宣言
func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()      //インスタンスを作成
	fmt.Println(err.Error()) //エラーを発生させる。 output -> カスタムエラーが発生しました
}
