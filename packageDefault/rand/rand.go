package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand
//擬似乱数を生成する

func main() {
	rand.Seed(256) //シードは固定すると出力結果が同じになるので注意。
	// シード値に基づいて乱数を生成するのでシードは流動的にした方が良い。

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	/* 実行するたび同じ値を出力する
	0.813527291469711
	0.5598026045235738
	0.6695717783859498
	*/

	// 現在時刻をシードに使った擬似乱数の生成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano()) //回避作として実行時刻を当てはめる
	// 0~99
	fmt.Println("0~99:", rand.Intn(100))
	// int
	fmt.Println("int:", rand.Int())
	// int32
	fmt.Println("int32:", rand.Int31())
	// int64
	fmt.Println("int64:", rand.Int63())
	// unit32
	fmt.Println("unit32:", rand.Uint32())
}
