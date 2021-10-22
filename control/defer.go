package main

import (
	"fmt"
	"os"
)

// defer
// 関数の終了時に実行されるもの

func TestDefer() {
	// defer fmt.Println("this is defer")
	fmt.Println("start")
	i := 3
	defer Test(i)
}

func Test(prop int) {
	for i := 0; i < prop; i++ {
		defer fmt.Println(i) //deferが複数ある場合、お尻から実行される
	}
}

func main() {
	TestDefer()

	// deferはファイルのcloseなどに使う
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello, React"))
}
