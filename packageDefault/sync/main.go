package main

import (
	"fmt"
	"sync"
)

// sync
// ミューテックスによる同期処理

// var st struct{ A, B, C int }

// これがだめな例 (同時に一つの変数にアクセスしようとするので衝突が起きてしまう)
// func UpdataAndPrint(n int) {
// 	st.A = n
// 	time.Sleep(time.Microsecond)
// 	st.B = n
// 	time.Sleep(time.Microsecond)
// 	st.C = n
// 	time.Sleep(time.Microsecond)
// 	fmt.Println(st)
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			for i := 0; i < 1000; i++ {
// 				UpdataAndPrint(i)
// 			}
// 		}()
// 	}

// 	for {

// 	}
// }

// // ミューテックスを定義
// var mutex *sync.Mutex

// func UpdateAndPrint(n int) {
// 	// ロックして他の関数がアクセスできないようにする(排他制御)
// 	mutex.Lock()

// 	st.A = n
// 	time.Sleep(time.Microsecond)
// 	st.B = n
// 	time.Sleep(time.Microsecond)
// 	st.C = n
// 	time.Sleep(time.Microsecond)
// 	fmt.Println(st)

// 	mutex.Unlock()
// }

// func main() {
// 	// mutexを初期化
// 	mutex = new(sync.Mutex)

// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			for i := 0; i < 1000; i++ {
// 				UpdateAndPrint(i)
// 			}
// 		}()
// 	}

// 	for {
// 	}

// }

// 全てのゴルーチンの終了を待ち受ける

func main() {
	// sync.WaitGroupを生成
	wg := new(sync.WaitGroup)
	// 待ち受ける個数を指定
	wg.Add(3)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st goroutine", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd goroutine", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd goroutine", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all done!")
}
