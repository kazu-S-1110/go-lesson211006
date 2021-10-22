package main

import "fmt"

// channnel
// 複数のgoroutinの間でデータの受け渡しをするために設計されたデータ構造
// 宣言と操作

func main() {
	var ch1 chan int //これだけだと読み取りも書き込みもできないので初期化が必要

	// 受信専用
	// var ch2 chan<- int

	// 送信専用
	// var ch3 <-chan int

	// 初期化
	ch1 = make(chan int)
	ch2 := make(chan int)
	fmt.Println(cap(ch1), cap(ch2))

	// バッファサイズ(限界容量、capacity)を指定して作成
	ch3 := make(chan int, 5)
	fmt.Println("ch3's capacity is", cap(ch3))

	// channelにデータを送信
	ch3 <- 100
	ch3 <- 200
	ch3 <- 300

	fmt.Println("ch3's length is", len(ch3))

	// channelからデータを受信(入れた順から取り出される)
	i := <-ch3
	fmt.Println("i is", i)
	fmt.Println("ch3's length is", len(ch3))
	i2 := <-ch3
	fmt.Println("i2 is", i2)
	fmt.Println("ch3's length is", len(ch3))
	// 直接受け取ることも可能
	fmt.Println(<-ch3)
	fmt.Println("ch3's length is", len(ch3))

	// バッファサイズを超えるとエラーを吐くので注意。
	ch3 <- 1
	ch3 <- 1
	ch3 <- 1
	ch3 <- 1
	ch3 <- 1
	// ch3 <- 1
	// fatal error: all goroutines are asleep - deadlock!
}
