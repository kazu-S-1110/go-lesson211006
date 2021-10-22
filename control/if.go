package main

import "fmt"

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 簡易if
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// 簡易ifの注意点
	if x := 2; true {
		fmt.Println("this is x" + string(x))
	}
	// fmt.Println(x) //error　簡易ifで定義した変数はそこでしか生きない。

}
