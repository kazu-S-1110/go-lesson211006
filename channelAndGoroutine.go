package main

import (
	"fmt"
	"time"
)

// channel
// channelとgoroutine間の受け渡し

func receiver(c chan int) {
	for {
		i := <-c
		fmt.Println("receive is", i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go receiver(ch1)
	go receiver(ch2)

	i := 0
	for i < 100 {
		if i%2 != 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
		time.Sleep(50 * time.Millisecond)
		i++
	}

}
