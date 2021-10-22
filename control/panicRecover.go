package main

import "fmt"

// panic, recover
// 意図的にruntimeerrorを発生させるもの、エラーハンドリングは別の方法が推奨されているのでここでは理解だけする

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("this is no") //表示されない
}
