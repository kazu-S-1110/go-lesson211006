package main

import (
	"fmt"
	"time"
)

// time

func main() {
	// 時間の生成
	// 今の時間
	t := time.Now()
	fmt.Println(t) //2021-10-26 07:53:59.614418 +0900 JST m=+0.000065210

	// 指定した時間を生成
	t2 := time.Date(2020, 1, 1, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay()) //通算日
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone()) //JST 32400

	// 時間の感覚を表現する
	// time.Duration型を返す
	fmt.Println(time.Hour)        //1h0m0s
	fmt.Printf("%T\n", time.Hour) //time.Duration
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond) //1µs
	fmt.Println(time.Nanosecond)  //1ns

	// time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("2h40m")
	fmt.Println(d) //2h40m0s

	t3 := time.Now()
	fmt.Printf("%T\n", t3)                      //time.Time
	fmt.Println(t3)                             //2021-10-26 08:08:25.461644 +0900 JST m=+0.001086042
	t3 = t3.Add(2*time.Minute + 15*time.Second) //2分15秒を足している(time.timeにtime.durationを組み合わせる)
	fmt.Println(t3)                             //2021-10-26 08:10:40.461644 +0900 JST m=+135.001086042

	// 時間の差分を取得
	t5 := time.Date(2021, 6, 24, 0, 0, 0, 0, time.Local)
	t6 := time.Now()
	fmt.Println(t6)

	// t5 - t6
	d2 := t5.Sub(t6)
	fmt.Println(d2) //-2984h16m16.33991s

	// 時刻を比較する
	fmt.Println(t6.Before(t5)) //false　t6がt5より前か判断してbool値で返す
	fmt.Println(t6.After(t5))  //true
	fmt.Println(t5.Before(t6))
	fmt.Println(t5.After(t6))

	// 指定時間のスリープ
	// 5秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("5秒止まってましたよ")
}
