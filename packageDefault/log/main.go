package main

import (
	"log"
	"os"
)

// log

func main() {
	// logの出力先を変更
	// log.SetOutput(os.Stdout)

	// log.Print("log\n")       //2021/10/27 07:52:12 log
	// log.Println("log2")      //2021/10/27 07:52:12 log2
	// log.Printf("log%d\n", 3) //2021/10/27 07:52:12 log3

	// log.Fatal("log\n") //logの日時とexit status 1が返される
	// log.Fatalln("log2")
	// log.Fatalf("log%d\n", 3)

	// log.Panic("log\n") //logの日時を出力して強制終了させる

	// 任意のファイルを作成して、logの出力先に指定
	// f, err := os.Create("test.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(f)
	// log.Println("ファイルに書き込む")

	log.SetOutput(os.Stdout)
	// ログのフォーマットを指定する
	// 標準のログフォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("a") //2021/10/27 08:08:40 a

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("b") //2021/10/27 08:08:40.704455 b

	// 時刻とファイルの行番号（短縮系）
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("c") //08:09:46 main.go:44: c

	log.SetFlags(log.LstdFlags)
	// ログのプリフィックスを設定
	log.SetPrefix("[LOG]")
	log.Println("d") //[LOG]2021/10/27 08:11:10 d

	// ロガーの生成
	// golangでは出力先やフォーマットを変えると全体に影響してしまう。一部だけログ出力を変更したいときはロガーを利用する
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("this is message for logger")
	log.Println("this is message for cmdline")

	// 条件分岐。ログを別に出力する
	// _, err := os.Open("fgaldkfjal")
	// if err != nil {
	// 	logger.Fatalln("Exit", err)
	// }
}
