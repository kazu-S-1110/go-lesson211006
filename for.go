package main

import "fmt"

// for

func main() {
	// for {
	// 	fmt.Println("loop") //延々と続いてしまう
	// }

	// 止める処理は何かしら入れないとだめ。breakなど
	// i := 0
	// for {
	// 	i++
	// 	if i == 34 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// whileみたいなやつ
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// 典型的なfor文
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue //スキップ
		}
		fmt.Println(i)
	}

	// 配列のfor
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// map配列(object)のfor
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
