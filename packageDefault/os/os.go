package main

import (
	"os"
)

func main() {
	// Exit
	// os.Exit(1)
	// fmt.Println("hello") //上のexitで終了するのでこのコードは実行されない

	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0) //deferも発火されないので注意。

	// log.Fatal //プログラムを強制終了させるようなもの？Printlnした後os.Exit(1)を実行する
	// _, err := os.Open("A.txt") //存在しないファイルを開こうとする
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Args
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// // os.Argsの要素数を表示
	// fmt.Printf("length=%d\n", len(os.Args))

	// // os.Argsの中身を全て表示
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }
	// // 実行方法　go buildしてbuildしたものを引数をつけて実行する

	// ファイル操作Open
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// ファイル操作Create
	f, _ := os.Create("hoge.txt") //createは新規作成するメソッド、既に存在する場合削除してもう一度作成する
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang\n"), 6) //6文字offsetしたところから書き込む
	f.Seek(0, os.SEEK_END)           //末尾に移動するメソッド
	f.WriteString("Yeah!")
}
