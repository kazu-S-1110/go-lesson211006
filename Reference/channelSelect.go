package main

import "fmt"

// channel
// select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 4)

	ch2 <- "A"

	ch2 <- "B"
	ch2 <- "C"
	ch2 <- "D"
	/*
		v1 := <-ch1 //ch1にはまだ何もないのでエラーを起こしその下のコードまで到達しない
		v2 := <-ch2 //これを制御するのがselect文
	*/

	// 実行されるものはランダムに行われる。（channelに入るものが複数の場合）
	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	case v3 := <-ch2:
		fmt.Println(v3 + "ok")
		// どれにも当てはまらない場合はdefaultに書く。
	default:
		fmt.Println("I don't know'")
	}

	// ch に値が入っている場合は case v := <- ch ケースが実行される。ch に値が入っていない場合は default ケースが実行される。default ケースが実行された場合はチャネルからの値の受信は行われない。
	// selectの活用例
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("received", i3)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100 {
		// 	break
		// }
	}
}
