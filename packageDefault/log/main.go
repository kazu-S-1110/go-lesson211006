package main

import (
	"log"
	"os"
)

// log

func main() {
	// logの出力先を変更
	// log.SetOutput(os.Stdout)

	log.Print("log\n")       //2021/10/27 07:52:12 log
	log.Println("log2")      //2021/10/27 07:52:12 log2
	log.Printf("log%d\n", 3) //2021/10/27 07:52:12 log3

	// log.Fatal("log\n") //logの日時とexit status 1が返される
	// log.Fatalln("log2")
	// log.Fatalf("log%d\n", 3)

	// log.Panic("log\n") //logの日時を出力して強制終了させる

	// 任意のファイルを作成して、logの出力先に指定
	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ファイルに書き込む")
}
