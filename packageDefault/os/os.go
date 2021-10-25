package main

import (
	"fmt"
	"log"
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
	// f, _ := os.Create("hoge.txt") //createは新規作成するメソッド、既に存在する場合削除してもう一度作成する
	// f.Write([]byte("Hello\n"))
	// f.WriteAt([]byte("Golang\n"), 6) //6文字offsetしたところから書き込む
	// f.Seek(0, os.SEEK_END)           //末尾に移動するメソッド
	// f.WriteString("Yeah!")

	// ファイル操作Read
	// f, err := os.Open("hoge.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// bs := make([]byte, 128) //第二引数はlength長さを指定
	// n, err := f.Read(bs)    //スライスを渡してそこに格納している
	// fmt.Println(n)          //nには書き込んだバイト数が入っている
	// fmt.Println(string(bs))

	// // readatを使うと移動した位置から読み取る
	// bs2 := make([]byte, 128)
	// nn, err := f.ReadAt(bs2, 10)
	// fmt.Println(nn)
	// fmt.Println(string(bs2))

	// ファイル操作OpenFile
	// O_RDONLY 読み取り専用
	// O_WRONLY 書き込み専用
	// O_RDWR 読み書き可能
	// O_APPEND ファイル末尾に追記
	// O_CREATE ファイルがなければ作成
	// O_TRUNC 可能であればファイルの内容をオープン時に空にする
	f, err := os.OpenFile("hoge.txt", os.O_RDWR|os.O_CREATE, 0666) //第二引数にはパイプでつなげることが可能（読み書きで開く、なければCreateする）
	// 第三引数はファイルを新規作成する際のパーミッションの値を入れる。開くとか書き込む際は意味がないお話、多分。
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
	f.Seek(0, os.SEEK_END)
	f.WriteString("okkkkkk")
}
