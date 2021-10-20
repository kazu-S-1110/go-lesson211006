package main

import "fmt"

// channel
// close

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 100
	ch1 <- 1000
	close(ch1)
	// closeしたchannelには送信は不可
	// ch1 <- 1 //panic: send on closed channel

	// 受信は可能
	fmt.Println(<-ch1)
	// channelの中身が0且つclose状態の判別
	i, ok := <-ch1
	fmt.Println(i, ok) //1000 true
	i2, ok := <-ch1
	fmt.Println(i2, ok) //0 false

}
