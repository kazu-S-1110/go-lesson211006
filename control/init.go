package main

import "fmt"

// init
// mainより先に実行される関数のこと
// JavaScriptでいうconstructor関数のことかな

func init() {
	fmt.Println("init")
}

// init関数は複数でもおk。実行順は上から。でも、分ける必要はない。
func init() {
	fmt.Println("init2")
}

func main() {
	// mainにinit関数を組み込んでいないがinit関数が先に実行される
	fmt.Println("Main")
}
