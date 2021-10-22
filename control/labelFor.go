package main

import "fmt"

// ラベル付きfor

func main() {
	// 下のコードのようなネストされたforを一気に抜けたい
	// for {
	// 	for {
	// 		for {
	// 			fmt.Println("START")
	// 			break
	// 		}
	// 		fmt.Println("noooo")
	// 	}
	// 	fmt.Println("noooo")
	// }
	// fmt.Println("END")

	// その時に使用するのがラベル付きfor
Loooooop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loooooop
			}
			fmt.Println("noooo") //表示されない
		}
		fmt.Println("noooo") //表示されない
	}
	fmt.Println("END")

Loooop:
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			// jが4以上の時は最初に戻るようにする
			if j == 4 {
				continue Loooop
			}
			fmt.Println(i, j, i*j)
		}
	}
}
