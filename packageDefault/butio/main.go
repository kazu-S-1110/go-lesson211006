package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio

func main() {
	// 標準入力を行単位で読み込む

	// 標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャンを成功する限り繰り返すループ
	for scanner.Scan() {
		fmt.Println("output->", scanner.Text())
	}

	// スキャンにエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
