package main

import (
	"fmt"
	"time"
)

// go文 goroutin
// 並行処理をかける。他の言語では難易度が高いがgoなら容易らしい。

func sub(str string) {
	for {
		fmt.Println("Sub Loop", str)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// sub() ←だけではsubが終わるまで下のコードは実行されない
	// goをつけると並行に処理してくれる
	go sub("hoge")
	go sub("sage")

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
