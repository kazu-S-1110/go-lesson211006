package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ioutil

func main() {
	// 入力全体を読み込む
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	// ファイルに書き込み(新規に作成される)（既存の内容は消される）
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
