package main

import (
	"fmt"
	"time"
)

// channel
// close

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "End")
}

func main() {
	ch1 := make(chan int, 2)

	// ch1 <- 100
	// ch1 <- 1000
	// close(ch1)
	// // closeしたchannelには送信は不可
	// // ch1 <- 1 //panic: send on closed channel

	// // 受信は可能
	// fmt.Println(<-ch1)
	// // channelの中身が0且つclose状態の判別
	// i, ok := <-ch1
	// fmt.Println(i, ok) //1000 true
	// i2, ok := <-ch1
	// fmt.Println(i2, ok) //0 false

	go receiver("No1:goroutine", ch1)
	go receiver("No2:goroutine", ch1)
	go receiver("No3:goroutine", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
	// 結果は毎回変わる。
	// No1:goroutine 1
	// No2:goroutine 2
	// No3:goroutine 0
	// No2:goroutine 3
}
